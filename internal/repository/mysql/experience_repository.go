package mysql

import "learning/personal/internal/entity"

type ExperienceRepository struct{}

func (r Repository) FetchAll() ([]entity.Experience, error) {
	return []entity.Experience{
		{
			ID:           1,
			Title:        "Web Developer",
			WorkLocation: "Stark Industries Los Angeles, CA",
			Content:      "Lorem ipsum dolor sit amet consectetur adipisicing elit. Delectus laudantium, voluptatem quis repellendus eaque sit animi illo ipsam amet officiis corporis sed aliquam non voluptate corrupti excepturi maxime porro fuga.",
		},
	}, nil
}
