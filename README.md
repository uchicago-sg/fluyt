# Fluyt

A faster, more portable backend for Marketplace. If you are hoping to submit
PRs to the live version, visit https://github.com/uchicago-sg/caravel.git.
This repository is primarily intended as a research prototype.

## Architecture

Fluyt is intended to handle a small corpus of data very, very quickly. 
Primarily, it does this by storing all listings in memory, and then streaming
them quickly from memcache, falling back to a persistent backend if necessary.

In the first prototype, we have an App Engine Datastore backend, as well as a
null backend. It would be relatively straightforward to add backends for 
AWS SimpleDB or PostgreSQL, since the database does not use native indexing.

## License

Copyright 2016 Jeremy Archer

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
