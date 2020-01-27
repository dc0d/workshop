# [Gossiping Bus Drivers Kata](https://kata-log.rocks/gossiping-bus-drivers-kata)

## Description

Bus drivers like to gossip, everyone knows that. And bus drivers can gossip when they end up at the same stop. So now we are going to calculate after how many stops all the bus drivers know all the gossips. You will be given a number of bus routes that the drivers follow. Each driver starts with one gossip. Each route is appointed to 1 driver. When 2 or more drivers are at the same stop (even if it is the start), they can exchange all the gossips they know. A route looks like this: 1 2 3 4 and is repeated over the whole day like this 1 2 3 4 1 2 3 4 1 2 3 … If a driver starts and stops at the same stop then that is also repeated. All drivers take 1 minute to go from one stop to another and the gossip exchange happens instantly. All drivers drive 8 hours a day so you have a maximum of 480 minutes to get all the gossiping around.

### Example 1:

Input:

```
3 1 2 3
3 2 3 1
4 2 3 4 5
```

Output:

```
5
```

### Example 2:

```
2 1 2
5 2 8
```

```
never
```
