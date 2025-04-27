package models

import "time"

type Page struct {
	ID                         uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title                      string     `gorm:"type:text;not null" json:"title"`
	Slug                       *string    `gorm:"type:text" json:"slug,omitempty"`
	MetaTags                   *string    `gorm:"column:meta_tags;type:text" json:"meta_tags,omitempty"`
	MetaDescription            *string    `gorm:"column:meta_description;type:text" json:"meta_description,omitempty"`
	Content                    *string    `gorm:"type:longtext" json:"content,omitempty"`
	Status                     *string    `gorm:"type:varchar(191)" json:"status,omitempty"`
	Visibility                 *string    `gorm:"type:varchar(191)" json:"visibility,omitempty"`
	PageBuilderStatus          *int16     `gorm:"column:page_builder_status;type:smallint" json:"page_builder_status,omitempty"`
	NavbarCategoryDropdownOpen *int16     `gorm:"column:navbar_category_dropdown_open;type:smallint" json:"navbar_category_dropdown_open,omitempty"`
	CreatedAt                  *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt                  *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	NavbarVariant              *string    `gorm:"column:navbar_variant;type:varchar(191)" json:"navbar_variant,omitempty"`
	BreadcrumbStatus           *string    `gorm:"column:breadcrumb_status;type:varchar(191)" json:"breadcrumb_status,omitempty"`
	PageContainerOption        *string    `gorm:"column:page_container_option;type:varchar(191)" json:"page_container_option,omitempty"`
}
