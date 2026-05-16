import time
import copy

class Twitter:

    def __init__(self):
        self.network = {}
        self.tweets = []

    def postTweet(self, userId: int, tweetId: int) -> None:
        timestamp = time.time()
        tweet = (-timestamp, userId, tweetId)
        self.tweets.append(tweet)

    def getNewsFeed(self, userId: int) -> List[int]:
        tweetsHeap = copy.deepcopy(self.tweets)
        heapq.heapify(tweetsHeap)
        if userId in self.network:
            follows = self.network[userId]
        else:
            follows = []
        i = 0
        ans = []
        print(tweetsHeap)
        print(follows)
        print('\n')
        while tweetsHeap and  i < 10:
            tweet = heapq.heappop(tweetsHeap)
            if tweet[1] == userId or tweet[1] in follows:
                ans.append(tweet[2])
                i += 1
        return ans

    def follow(self, followerId: int, followeeId: int) -> None:
        if followerId not in self.network.keys():
            self.network[followerId] = [followeeId]
        elif followeeId not in self.network[followerId]:
            self.network[followerId].append(followeeId)

    def unfollow(self, followerId: int, followeeId: int) -> None:
        if followerId in self.network:
            self.network[followerId].remove(followeeId)
            print(self.network)
            if len(self.network[followerId]) == 0:
                del self.network[followerId]
