package loader

import (
	"fmt"
	"math/rand"
	"time"
)

type QAPair struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Answer string `json:"answer"`
	Tags   string `json:"tags"`
	Score  int    `json:"score"`
}

var (
	programmingTopics = []string{
		"JavaScript", "Python", "Go", "Rust", "Java", "C++", "TypeScript",
		"React", "Node.js", "PostgreSQL", "Docker", "Kubernetes", "AWS",
		"Git", "Linux", "API", "REST", "GraphQL", "MongoDB", "Redis",
	}

	questionStarters = []string{
		"How to", "Best way to", "Why does", "What is the",
		"Cannot get", "Help with", "Issue with", "Problem running",
		"Understanding", "Difference between", "When should I use",
		"Is it possible to", "How do I implement", "Need help optimizing",
	}

	topicActions = map[string][]string{
		"JavaScript": {"handle async/await", "access object properties", "iterate over arrays", "use closures", "manage state"},
		"Python":     {"use list comprehensions", "handle exceptions", "create virtual environments", "import modules", "use decorators"},
		"Go":         {"write concurrent code", "handle errors properly", "use interfaces", "manage goroutines", "structure projects"},
		"PostgreSQL": {"optimize queries", "use indexes", "handle transactions", "create views", "use JSON columns"},
		"Docker":     {"build images", "manage volumes", "use docker-compose", "optimize Dockerfiles", "network containers"},
		"API":        {"design REST endpoints", "handle authentication", "version APIs", "document APIs", "use middleware"},
	}

	answerTemplates = []string{
		"You can achieve this by using the standard library. Here's an example:\n\n```\ncode example here\n```\n\nThis approach is recommended because it's simple and efficient.",
		"The best practice is to follow these steps:\n1. First, do this\n2. Then, do that\n3. Finally, handle the result\n\nMake sure to handle edge cases properly.",
		"I had the same issue. The solution is to use a different approach:\n\n```\ncode\n```\n\nThis worked for me in production with high load.",
		"Actually, you should consider using a library like X or Y. They handle this internally and are well-maintained.\n\n```\nexample\n```\n\nThis gives you better performance.",
		"The problem might be in your configuration. Try setting these values:\n\n```\nconfig\n```\n\nIf that doesn't work, check the logs for specific errors.",
	}
)

func GenerateQA(count int) []QAPair {
	rand.Seed(time.Now().UnixNano())

	pairs := make([]QAPair, count)

	for i := 0; i < count; i++ {
		topic := programmingTopics[rand.Intn(len(programmingTopics))]
		starter := questionStarters[rand.Intn(len(questionStarters))]

		var action string
		if actions, ok := topicActions[topic]; ok {
			action = actions[rand.Intn(len(actions))]
		} else {
			action = "implement this feature"
		}

		pairs[i] = QAPair{
			ID:     int64(i + 1),
			Title:  fmt.Sprintf("%s %s in %s", starter, action, topic),
			Body:   generateQABody(topic, action),
			Answer: generateAnswer(topic),
			Tags:   generateTags(topic),
			Score:  rand.Intn(100) - 20,
		}
	}

	return pairs
}

func generateQABody(topic, action string) string {
	templates := []string{
		"I'm trying to %s in %s but I'm running into issues. I've tried several approaches but none seem to work correctly.\n\nHere's what I have so far:\n\n```\ncode snippet\n```\n\nWhat am I doing wrong? Any help would be appreciated.",
		"New to %s and need help with %s. I've read the documentation but it's not clear how to handle edge cases.\n\nCan someone explain the best practices for this scenario?",
		"Looking for the most efficient way to %s with %s. Performance is important as this will run frequently.\n\nCurrent implementation is too slow for production use.",
	}

	template := templates[rand.Intn(len(templates))]
	return fmt.Sprintf(template, action, topic)
}

func generateAnswer(topic string) string {
	template := answerTemplates[rand.Intn(len(answerTemplates))]

	examples := map[string]string{
		"JavaScript": "const result = await fetch('/api/data');",
		"Python":     "result = await fetch_data()",
		"Go":         "result, err := fetchData()",
		"PostgreSQL": "SELECT * FROM users WHERE active = true",
		"Docker":     "docker build -t myapp .",
	}

	code := "your code here"
	if ex, ok := examples[topic]; ok {
		code = ex
	}

	return fmt.Sprintf(template, code)
}

func generateTags(topic string) string {
	tagSets := map[string][]string{
		"JavaScript": {"javascript", "web", "frontend"},
		"Python":     {"python", "backend", "scripting"},
		"Go":         {"go", "golang", "backend"},
		"PostgreSQL": {"postgresql", "database", "sql"},
		"Docker":     {"docker", "devops", "containers"},
		"API":        {"api", "rest", "backend"},
	}

	if tags, ok := tagSets[topic]; ok {
		return tags[0] + "," + tags[1]
	}
	return topic + ",programming"
}
