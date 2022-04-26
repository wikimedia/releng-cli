#!/usr/bin/env bash
#
# Tests general docker commands

set -e # Fail on errors
SCRIPT_DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
source $SCRIPT_DIR/functions.sh
source $SCRIPT_DIR/pretest-mediawiki.sh

export MWCLI_CONTEXT_TEST=1

function finish {
    cd $SCRIPT_DIR/..

    # Show it all
    docker ps

    # Destroy it all
    test_command_success "./bin/mw docker destroy --no-interaction"

    # Clean up & make sure no services are running
    test_docker_ps_service_count 0
    if ./bin/mw docker hosts writable --no-interaction; then
        test_command_success "./bin/mw docker hosts remove --no-interaction"
    else
        echo "sudo needed for hosts file modification!"
        test_command_success "sudo -E ./bin/mw docker hosts remove --no-interaction"
    fi
    test_command_success "./bin/mw docker env clear --no-interaction"
}
trap finish EXIT

test_command_success "./bin/mw docker env clear --no-interaction"

# Run this integration test using a non standard port, unlikley to conflict, to make sure it works
test_command_success "./bin/mw docker env set PORT 6194"
# And already fill in the location of mediawiki
MWDIR=$(pwd)/.mediawiki
test_command_success "./bin/mw docker env set MEDIAWIKI_VOLUMES_CODE ${MWDIR}"

# Setup the default hosts in hosts file
if ./bin/mw docker hosts writable --no-interaction; then
    test_command_success "./bin/mw docker hosts add --no-interaction"
else
    echo "sudo needed for hosts file modification!"
    test_command_success "sudo -E ./bin/mw docker hosts add --no-interaction"
fi

# Create
test_command_success "./bin/mw docker mediawiki create"

# Get the port in use
PORT=$(./bin/mw docker env get PORT)

# Make sure that exec generally works as expected
./bin/mw docker mediawiki exec -- FOO=bar env | grep FOO

# Validate the basic stuff
test_command_success "./bin/mw docker docker-compose ps"
test_command_success "./bin/mw docker env list"

test_curl http://default.mediawiki.mwdd.localhost:$PORT "Could not find a running database for the database name"

# Install sqlite & check
test_command_success "./bin/mw docker mediawiki install --dbtype sqlite"
test_curl http://default.mediawiki.mwdd.localhost:$PORT "MediaWiki has been installed"

# cd to mediawiki
cd .mediawiki

# composer: Make sure a command works in root of the repo
test_command "./../bin/mw docker mediawiki composer home" "https://www.mediawiki.org/"

# exec: Make sure a command works in the root of the repo
test_command "./../bin/mw docker mediawiki exec ls" "api.php"

# exec phpunit: Make sure using exec to run phpunit things works
test_command "./../bin/mw docker mediawiki exec -- composer phpunit tests/phpunit/unit/includes/PingbackTest.php" "OK "

# fresh: Make sue a basic browser test works
test_command_success "./../bin/mw docker mediawiki fresh npm run selenium-test -- -- --spec tests/selenium/specs/page.js"

# quibble: Make sure a quibble works
test_command_success "./../bin/mw docker mediawiki quibble quibble -- --help"
test_command "./../bin/mw docker mediawiki quibble quibble -- --skip-zuul --skip-deps --skip-install --db-is-external --command \"ls\"" "index.php"

# cd to Vector
cd skins/Vector

# composer: Make sure a command works from the Vector directory
test_command "./../../../bin/mw docker mediawiki composer home" "http://gerrit.wikimedia.org/g/mediawiki/skins/Vector"
# exec: Make sure a command works from the Vector directory
test_command "./../../../bin/mw docker mediawiki exec ls" "skin.json"

# gerrit project current
test_command "./../../../bin/mw gerrit project current" "mediawiki/skins/Vector"
