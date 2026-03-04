package loader

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Article struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var (
	wikiSubjects = []string{
		"History", "Science", "Geography", "Arts", "Technology",
		"Sports", "Medicine", "Music", "Literature", "Philosophy",
		"Politics", "Economics", "Biology", "Chemistry", "Physics",
		"Astronomy", "Mathematics", "Psychology", "Sociology", "Law",
	}

	wikiPrefixes = []string{
		"Introduction to", "The History of", "A Guide to", "Understanding",
		"Fundamentals of", "Advanced", "Basic", "Modern", "Classical",
		"The Nature of", "Exploring", "Analysis of", "Principles of",
	}

	wikiTopics = map[string][]string{
		"History":    {"World War II", "Ancient Rome", "The Renaissance", "Industrial Revolution", "French Revolution", "Cold War", "Medieval Europe", "Roman Empire", "Greek Civilization", "American Civil War"},
		"Science":    {"Quantum Mechanics", "Evolution Theory", "Big Bang", "Atomic Structure", "Cell Biology", "Genetics", "Thermodynamics", "Electromagnetism", "Relativity", "Neuroscience"},
		"Geography":  {"Mount Everest", "Amazon Rainforest", "Pacific Ocean", "Sahara Desert", "Himalayas", "Great Barrier Reef", "Nile River", "Alps", "Grand Canyon", "Antarctica"},
		"Technology": {"Artificial Intelligence", "Internet of Things", "Blockchain", "Quantum Computing", "5G Networks", "Cloud Computing", "Machine Learning", "Cybersecurity", "Virtual Reality", "Renewable Energy"},
		"Medicine":   {"COVID-19", "Cancer Treatment", "DNA Sequencing", "Vaccines", "Antibiotics", "Organ Transplants", "Mental Health", "Heart Disease", "Diabetes", "Alzheimer's Disease"},
		"Arts":       {"Renaissance Art", "Impressionism", "Modern Art", "Sculpture", "Architecture", "Photography", "Cinema", "Dance", "Theater", "Pottery"},
		"Music":      {"Classical Music", "Jazz", "Rock and Roll", "Hip Hop", "Electronic Music", "Folk Music", "Blues", "Country Music", "Opera", "Reggae"},
		"Literature": {"Shakespeare", "Greek Mythology", "The Bible", "Harry Potter", "Don Quixote", "War and Peace", "Ulysses", "Moby Dick", "Pride and Prejudice", "1984"},
		"Sports":     {"Football", "Basketball", "Tennis", "Swimming", "Marathon", "Olympics", "Cricket", "Baseball", "Golf", "Boxing"},
		"Biology":    {"DNA", "Cells", "Photosynthesis", "Evolution", "Ecosystems", "Viruses", "Bacteria", "Animals", "Plants", "Human Body"},
	}

	wikiContentTemplates = []string{
		"%s is a significant topic in the field of %s. It has been studied extensively and continues to be an area of active research and discussion. The subject encompasses various aspects that are important to understand for both academic and practical purposes.\n\nHistorical background reveals that %s has evolved considerably over time. Early documentation dates back centuries, and our understanding has grown exponentially since then. Modern perspectives have shifted from traditional viewpoints to more nuanced interpretations.\n\nKey concepts include fundamental principles that govern %s. These include detailed mechanisms and processes that interact in complex ways. Understanding these relationships is essential for anyone looking to master this subject.\n\nApplications of this knowledge are widespread across numerous industries and disciplines. From practical implementations to theoretical frameworks, %s plays a crucial role in advancing our capabilities and understanding of the world.",
		"The study of %s has produced remarkable insights into the nature of %s. Researchers have dedicated countless hours to unraveling its mysteries, leading to groundbreaking discoveries that have shaped our modern world.\n\nOne of the most fascinating aspects of %s is its intricate relationship with other phenomena. This interconnection creates a rich tapestry of cause and effect that scientists continue to explore. The implications extend far beyond initial expectations.\n\nPractical implementations have demonstrated the value of understanding %s. Industries ranging from healthcare to technology have benefited from these advances. The future holds even more promise as new methodologies and tools emerge.\n\nControversies and debates surround certain aspects of %s. Different perspectives offer valuable insights, and ongoing discussions help refine our understanding. This dynamic environment ensures continued growth and learning.",
	}
)

func GenerateWikipediaArticles(count int) []Article {
	rand.Seed(time.Now().UnixNano())

	articles := make([]Article, count)

	for i := 0; i < count; i++ {
		subject := wikiSubjects[rand.Intn(len(wikiSubjects))]
		topic := wikiTopics[subject][rand.Intn(len(wikiTopics[subject]))]
		prefix := wikiPrefixes[rand.Intn(len(wikiPrefixes))]

		articles[i] = Article{
			ID:      int64(i + 1),
			Title:   fmt.Sprintf("%s %s", prefix, topic),
			Content: generateWikiContent(topic, subject),
		}
	}

	return articles
}

func generateWikiContent(topic, subject string) string {
	template := wikiContentTemplates[rand.Intn(len(wikiContentTemplates))]

	topicLower := strings.ToLower(topic)
	subjectLower := strings.ToLower(subject)

	content := fmt.Sprintf(template, topic, subject, topic, topic, topic, topic)
	content = strings.ReplaceAll(content, "the the", "the")

	_ = topicLower
	return content + fmt.Sprintf("\n\nSee also: Related topics in %s", subjectLower)
}
