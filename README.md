# A TDD Workflow; TCR

Source: [Parrot Refactoring Kata](https://github.com/emilybache/Parrot-Refactoring-Kata)

Considering the behaviour is fully described inside `workshop_test.go`, there is no need to have another set of tests for newly introduced `typeEuropean`, `typeAfrican` and `typeNorwegianBlue`. The relation between code and test is not per-class or per-struct. It's per-behaviour.