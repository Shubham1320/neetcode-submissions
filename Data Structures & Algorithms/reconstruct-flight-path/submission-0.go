func findItinerary(tickets [][]string) []string {

    adj := make(map[string][]string)
	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		adj[src] = append(adj[src], dst)
	}
	for src := range adj {
		sort.Sort(sort.Reverse(sort.StringSlice(adj[src])))
	}

	stack := []string{"JFK"}
	var res []string

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		if len(adj[curr]) == 0 {
			res = append(res, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, adj[curr][len(adj[curr])-1])
			adj[curr] = adj[curr][:len(adj[curr])-1]
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res


}
