# dicetrace

[![cicd](https://github.com/ngoldack/dicetrace/actions/workflows/cicd.yaml/badge.svg)](https://github.com/ngoldack/dicetrace/actions/workflows/cicd.yaml)
[![codecov](https://codecov.io/github/ngoldack/dicetrace/branch/main/graph/badge.svg?token=4IljTKeRUs)](https://codecov.io/github/ngoldack/dicetrace)
[![wakatime](https://wakatime.com/badge/github/ngoldack/dicetrace.svg)](https://wakatime.com/badge/github/ngoldack/dicetrace)

Monorepo containing the sourcecode for [dicetrace.io](https://dicetrace.io).

## Architecture

For more information visit the [documentation](https://docs.dicetrace.io/).

### Services

dicetrace consists of these services:

| Name           | Status                                                                                 | Function                                                                       | Language / Framework                                                                                                                                                                                                                        | Additional Information |
|----------------|----------------------------------------------------------------------------------------|--------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------|
| api            | ![progress](https://img.shields.io/badge/PROGRESS---?style=for-the-badge&color=yellow) | service for responsible for communicating between the ui and internal services | ![Go](https://img.shields.io/badge/go-00ADD8.svg?&style=for-the-badge&logo=go&logoColor=white)                                                                                                                                              |                        |
| user           | ![progress](https://img.shields.io/badge/PROGRESS---?style=for-the-badge&color=yellow) | service for managing user data                                                 | ![Go](https://img.shields.io/badge/go-00ADD8.svg?&style=for-the-badge&logo=go&logoColor=white)                                                                                                                                              |                        |
| collection     | ![progress](https://img.shields.io/badge/PROGRESS---?style=for-the-badge&color=yellow) | service for managing games and collections                                     | ![Go](https://img.shields.io/badge/go-00ADD8.svg?&style=for-the-badge&logo=go&logoColor=white)                                                                                                                                              |                        |
| group          | ![progress](https://img.shields.io/badge/PROGRESS---?style=for-the-badge&color=yellow) | service for managing groups                                                    | ![Go](https://img.shields.io/badge/go-00ADD8.svg?&style=for-the-badge&logo=go&logoColor=white)                                                                                                                                              |                        |
| match          | ![planned](https://img.shields.io/badge/PLANNED---?style=for-the-badge&color=red)      | service for realtime matches                                                   | ![Rust](https://img.shields.io/badge/rust-000000.svg?&style=for-the-badge&logo=rust&logoColor=white)                                                                                                                                              |                        |
| chat           | ![planned](https://img.shields.io/badge/PLANNED---?style=for-the-badge&color=red)      | service for realtime chat                                                      | ![Rust](https://img.shields.io/badge/rust-000000.svg?&style=for-the-badge&logo=rust&logoColor=white)                                                                                                                                                                                                                                           |                        |
| web            | ![progress](https://img.shields.io/badge/PROGRESS---?style=for-the-badge&color=yellow) | web app                                                                        | ![TypeScript](https://img.shields.io/badge/typescript-3178C6.svg?&style=for-the-badge&logo=typescript&logoColor=white)<br/>![SvelteKit](https://img.shields.io/badge/sveltekit-FF3E00.svg?&style=for-the-badge&logo=svelte&logoColor=white) |                        |
| desktop        | ![planned](https://img.shields.io/badge/PLANNED---?style=for-the-badge&color=red)      | desktop app                                                                    | ![Rust](https://img.shields.io/badge/rust-000000.svg?&style=for-the-badge&logo=rust&logoColor=white)<br />![Tauri](https://img.shields.io/badge/tauri-ffc131.svg?&style=for-the-badge&logo=tauri&logoColor=white)<br />![TypeScript](https://img.shields.io/badge/typescript-3178C6.svg?&style=for-the-badge&logo=typescript&logoColor=white)<br/>![Svelte](https://img.shields.io/badge/svelte-FF3E00.svg?&style=for-the-badge&logo=svelte&logoColor=white)                                                                                                                                                                                                                                           |
| mobile         | ![planned](https://img.shields.io/badge/PLANNED---?style=for-the-badge&color=red)      | mobile app                                                                     | ![Flutter](https://img.shields.io/badge/flutter-02569b.svg?&style=for-the-badge&logo=flutter&logoColor=white)                                                                                                                                                                                                                                           |                        |
| recommendation | ![planned](https://img.shields.io/badge/PLANNED---?style=for-the-badge&color=red)      | service for getting recommendations                                            | ![Python](https://img.shields.io/badge/python-3776ab.svg?&style=for-the-badge&logo=python&logoColor=white)<br />![TensorFlow](https://img.shields.io/badge/tensorflow-ff6f00.svg?&style=for-the-badge&logo=tensorflow&logoColor=white)                                                                                                                                                                                                                                           |                        |
| status         | ![progress](https://img.shields.io/badge/PROGRESS---?style=for-the-badge&color=yellow) | status website for all services                                                | ![uptime-kuma](https://img.shields.io/badge/uptime_kuma-5cdd8b.svg?&style=for-the-badge&logo=uptimekuma&logoColor=white)                                                                                                                    |                        |
| docs           | ![done](https://img.shields.io/badge/DONE---?style=for-the-badge&color=green)          | documentation for all services and their apis                                  | ![TypeScript](https://img.shields.io/badge/typescript-3178C6.svg?&style=for-the-badge&logo=typescript&logoColor=white)<br/>![Docusaurus](https://img.shields.io/badge/Docusaurus-3ecc5f.svg?&style=for-the-badge)                           |                        |

## Tech stack

| Name                                                                                                                | Type                     | Function                                                            |
|---------------------------------------------------------------------------------------------------------------------|--------------------------|---------------------------------------------------------------------|
| ![nx](https://img.shields.io/badge/nx-143055.svg?&style=for-the-badge&logo=nx&logoColor=white) | Build Tool               | build tool to build and manage this repo                            |
| ![terraform](https://img.shields.io/badge/terraform-7b42bc.svg?&style=for-the-badge&logo=terraform&logoColor=white) | IaC Tool                 | iac tool to deploy all services                                     |
| ![redis](https://img.shields.io/badge/redis-%23DD0031.svg?&style=for-the-badge&logo=redis&logoColor=white)          | Database                 | cache for the [api](#architecture#services) service                 |
| ![mongodb](https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb&logoColor=white)            | Database                 | database for the [user]() service                                   |
| ![neo4j](https://img.shields.io/badge/Neo4j-018bff?style=for-the-badge&logo=neo4j&logoColor=white)                  | Database                 | database for the [collection]() service                             |
| ![kafka](https://img.shields.io/badge/Kafka-231F20?style=for-the-badge&logo=apachekafka&logoColor=white)            | Event Streaming Platform | event streaming platform for communication of all internal services |
| ![postgresql](https://img.shields.io/badge/postgresql-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)   | Database                 | database for the [chat]() service                                   |
