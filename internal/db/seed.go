package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/Steve-Rynjah/go-social-feeds/internal/store"
)

var usernames = []string{
	"head",
}

var titles = []string{
	"Morning Glow", "Silent Waves", "Hidden Path", "Golden Hour", "Lost Echo",
	"Dream Chaser", "Frozen Time", "Crimson Sky", "Moon Whisper", "Shattered Light",
	"Beyond Stars", "Endless Journey", "Echoes Within", "Fading Memories", "Broken Truth",
	"Whispered Secrets", "Lonely Horizon", "Velvet Night", "Burning Dreams", "Timeless Soul",
}

var contents = []string{
	"A calm sunrise fills the sky with soft golden hues.",
	"The ocean whispers stories only the moon can hear.",
	"A forgotten trail leads to places untouched by time.",
	"Sunlight paints the city in shades of hope.",
	"Echoes of laughter fade into the distant hills.",
	"Chasing dreams that dance just beyond reach.",
	"Moments frozen in the chill of yesterday’s breeze.",
	"A crimson sunset melts into the horizon.",
	"The moon hums softly above sleeping clouds.",
	"Light shatters across the glass of dawn.",
	"Stars guide those who dare to wander far.",
	"The road stretches into infinity, waiting to be walked.",
	"Every echo tells a story of the soul within.",
	"Memories fade, but the feeling lingers on.",
	"The truth breaks, leaving shards of realization.",
	"Secrets hide between the words never spoken.",
	"The horizon stands still, watching the lone traveler.",
	"The night wraps the world in velvet silence.",
	"Dreams burn bright, even in the darkest hours.",
	"A soul untouched by time finds peace in stillness.",
}

var tags = []string{
	"dev", "ui", "ux", "api", "db", "app", "web", "ios", "ml", "ai",
	"bot", "vr", "ar", "qa", "ops", "api2", "sys", "devx", "cli", "net",
}

var comments = []string{
	"Nice work!",
	"Looks great!",
	"Awesome job!",
	"Well done!",
	"Perfect!",
	"Good effort!",
	"Love this!",
	"So cool!",
	"Impressive!",
	"Fantastic!",
	"Great idea!",
	"Keep it up!",
	"Very nice!",
	"Excellent!",
	"Brilliant!",
	"Super clean!",
	"Nicely done!",
	"Cool update!",
	"Smooth work!",
	"Top notch!",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(1)

	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creation in user : ", err)
			return
		}
	}

	posts := generatePosts(200, users)

	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creation in post : ", err)
			return
		}
	}

	comments := generateComments(200, users, posts)

	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creation in comment : ", err)
			return
		}
	}

	log.Println("Seeding completed.")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "123456",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}

	return cms
}
