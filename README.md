# Gilded Rose Kata

Source: [Gilded Rose Kata](https://github.com/emilybache/GildedRose-Refactoring-Kata)

> The goal here is to practice a technique for refactoring legacy code.

# Steps Taken; A High-Level View

There were two phases:
- Phase 1: Have a test with 100% coverage; to preserve the bhaviour of the current system. After reaching 100% coverage, we will not touch this test and it acts as the measure for evaluating the next steps.
- Phase 2: Refactoring the code.

### Phase 1

- Adding different inputs and crystalize the expected output by running the test and getting the expected transformation on the item.
- At each step, it is possible to look into code for constant values that affect the logic (those magic strings and numbers), to have a pair of (input, expected-output) that increases the coverage a bit further.

### Phase 2

- Here the goal if to refactor this code and put it in a more meaningful structure.
- As we see there are items and each item has a name, so we can dispatch the logic based on the item's name.
- One technique is having a switch statement based on the item's name and copy the whole body of the function to each branch.
- After having the branches, with a green test, it can be seen that the coverage drops. It's because we branched based on the name of the item and some of the logic there is working no more.
- After having 100% coverage again, it's possible to replace that conditional statement (here, the switch statement) with polymorphism. For example here a `qualifier` interface is defined and implemented accordingly for each item.