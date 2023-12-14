# Tackle CLI tool

A tool for Konveyor Tackle application maintenance written in Python.

For more details about the Tackle project, see [Tackle2-Hub README](https://github.com/konveyor/tackle2-hub) or https://github.com/konveyor/tackle-documentation.

## Scenarios

### Export and Import data between Tackle/Konveyor instances

Tackle 2 can be exported and imported to another Tackle 2 instance, expecting it is cleaned-up before import. This focuses on application inventory, but exports and imports most of Tackle 2 objects except TaskGroups&Tasks and Keycloak identitty provider data.

Export/imported entities:
- Applications, Dependencies, Assessments, Reviews, Tag Types, Tags, Job Functions, Stakeholder Groups, Stakeholders, Business Services, Proxies, Identities (sensitive data are encrypted with ```encryption_passphase``` config file field)
- Application's Buckets data (e.g. app binaries or windup reports)

Steps for export (tackle-config.yml file points to source Tackle 2 installation)
- ```tackle export```

Steps for import (tackle-config.yml file points to destination Tackle 2, existing data will be removed)

- ```tackle clean-all```
- ```tackle import```

**Note: This export/import functionality doesn't aim to be a backup/restore tool. E.g. it doesn't dump Keycloak data and it requires have the Tackle installation to be running on the destination cluster before importing the data.**

### If the import failed

The ```tackle import``` command could fail in a pre-import check phase which ensures there are no resources of given type with the same ID in the destination Tackle 2 (error after ```Checking tagtypes in destination Tackle2..``` etc.). In this case, run ```tackle clean``` command, which will remove such objects from the destination Tackle 2 API or remove it manually either from the destination Tackle 2 or from the JSON data files.

If the import failed in the upload phase (error after ```Uploading tagtypes..``` etc.), try  ```tackle clean``` to remove already imported objects followed by  ```tackle clean-all``` which lists all resources of all known data types in the destination Tackle 2 API and deletes it (without looking to local data files).

Note on ```clean-all``` command, it deletes all resources from Tackle 2 Hub API, however Pathfinder API doesn't support listing assessments without providing an applicationID. The applications could not be present in Hub, so an "orphaned" assessments could stay in Pathfinder. In order to resolve potential collision with imported data, run  ```tackle clean``` together with the ```clean-all``` command.

Caution: all clean actions might delete objects already present in the Tackle 2 and unrelated to the import data.

## Requirements

The tool requires Python3 with YAML parser and PyCrypto package (or its successors like pycroyptodome) to be installed. A Python 3.6 (default in RHEL8) has YAML already included, but e.g. Python 3.9 (default in RHEL9) requires install a PyYAML module. Also git is needed to get the source code.

Install system requirements with  ```dnf install python3 python3-pip git``` for RHEL-like Linux or corresponding for your operating system.

Then install required Python libraries PyYAML and pycryptodome with PIP tool ```python3 -m pip install pyyaml pycryptodome```. Note: Install ```pycryptodome``` only if there is no library providing PyCrypto features already present on your system.

## Usage

Clone Github repository:
```git clone https://github.com/konveyor/tackle2-hub.git```

Change to the tool directory:
```cd hack/tool```

Use ```tackle-config.yml.example``` file as a template to set your Tackle endpoints and credentials and save it as ```tackle-config.yml```.

Run the tackle tool:
```./tackle```

### Supported actions
- ```export``` exports Tackle 2 objects into local JSON files and bucket data directory
- ```import``` creates objects in Tackle 2 from local JSON files and buckets data
- ```clean``` deletes objects uploaded to Tackle 2 from local JSON files
- ```clean-all``` deletes ALL data returned by Tackle 2 (including seeds, additional to ```clean```), skips some pathfinder stuff without index action

### Export Tackle 2

Run ```tackle export``` to get dump of Tackle 2 objects into JSON files in local directory ```./tackle-data``` and buckets data dump to ```./tackle-data/buckets``` directory.

The export dommand connects to Tackle2 Hub API and dumps relevant resources to local JSON files. For credentials/identities resources, their sensitive fields are encrypted with ```encryption_passphase``` from config file. Application's bucket data are gathered from API and stored locally as .tar.gz archives.

**Note: the encryption_passphase needs to be the same when running the import.**

### Import to Tackle 2

Check local JSON dump files in ```./tackle-data``` directory (if needed) and create objects in Tackle 2 Hub running ```tackle import```.

The import command connects to Tackle2 Hub, check existing objects for possible collisions (by IDs) and uploads resources from local JSON files.

### Delete uploaded objects

To delete objects previously created by the ```import``` command, run ```tackle clean```. This can address also existing Tackle 2 objects which are in collision with local JSON dump files.

### Delete all objects

The Tackle2 instance could be cleaned-up with ```tackle clean``` command. It lists objects from all data types and deletes such resources.

There is a exception with Pathfinder API which doesn't support listing assessments without knowledge of applicationIDs, so this might stay in Pathfinder database.

### Command line options

Config file ```-c / --config``` path specifies a YAML file with configuration options including endpoints and credentials for Tackle APIs (```tackle-config.yml``` by default).

Verbose output ```-v / --verbose``` option logs all API requests and responses providing more information for possible debugging (```off``` by default).

Data directory ```-d / --data-dir``` specifies path to local directory with Tackle database records in JSON format (```./tackle-data``` by default).

A full export without having access to the destination Tackle 2 and including all seed objects can be executed with ```-s / --skip-destination-check``` option. When importing such data, the destination Tackle 2 needs to be empty (run ```clean-all``` first). This is ```on``` by default in Tackle 2 ```export``` command.

SSL warnings ```-w / --disable-ssl-warnings``` optional suppress ssl warning for api requests.

Import errors could be skipped with ``` -i / --ignore-import-errors ``` -  not recommended - use with high attention to avoid data inconsistency. If the import has failed, it is recommended use ```tackle clean``` command to remove only imported resources.

Tackle2 deployments without auth feature (```feature_auth_required: false```), should use ```-n / --no-auth``` flag to skip Keycloak auth token creation, using empty token for Tackle API calls.

Tackle 2 export dumps buckets content which could be large, to skip it, use ```-b / --skip-buckets``` flag on export command.

## Example

```
usage: tackle [-h] [-c [CONFIG]] [-d [DATA_DIR]] [-v] [-s] [-w] [-i] [-n] [-b] [action ...]

Konveyor Tackle maintenance tool.

positional arguments:
  action                One or more Tackle commands that should be executed, options: export import clean clean-all

options:
  -h, --help            show this help message and exit
  -c [CONFIG], --config [CONFIG]
                        A config file path (tackle-config.yml by default).
  -d [DATA_DIR], --data-dir [DATA_DIR]
                        Local Tackle data directory path (tackle-data by default).
  -v, --verbose         Print verbose output (including all API requests).
  -s, --skip-destination-check
                        Skip connection and data check of Tackle 2 destination.
  -w, --disable-ssl-warnings
                        Do not display warnings during ssl check for api requests.
  -i, --ignore-import-errors
                        Skip to next item if an item fails load.
  -n, --no-auth         Skip Keycloak token creation, use empty Auth token in Tackle API calls.
  -b, --skip-buckets    Skip Tackle 2 Buckets content export.
```

API endpoints and credentials should be set in a config file (```tackle-config.yml``` by default).

```
---
# Main Tackle 2 endpoint and credentials
url: https://tackle-konveyor-tackle.apps.cluster.local
username: admin
password:

# Export of Identitiy (credentials) password and key fields should be encrypted, set the passphase
encryption_passphase:

```

Unverified HTTPS warnings from Python could be hidden by ```export PYTHONWARNINGS="ignore:Unverified HTTPS request"``` or with ```-w``` command option.
