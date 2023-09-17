from collections import defaultdict

from typing import List


class Solution:
    def findItinerary(self, tickets: List[List[str]]) -> List[str]:
        graph = defaultdict(list) 

        for src, dst in sorted(tickets, reverse=True):
            graph[src].append(dst)

        itinerary: List[str] = []
        def dfs(airport: str):
            while graph[airport]:
                dfs(graph[airport].pop())
            itinerary.append(airport)

        dfs("JFK")

        return itinerary[::-1]