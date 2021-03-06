// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "price", Type: field.TypeFloat64},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// ProductUsersColumns holds the columns for the "product_users" table.
	ProductUsersColumns = []*schema.Column{
		{Name: "product_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// ProductUsersTable holds the schema information for the "product_users" table.
	ProductUsersTable = &schema.Table{
		Name:       "product_users",
		Columns:    ProductUsersColumns,
		PrimaryKey: []*schema.Column{ProductUsersColumns[0], ProductUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "product_users_product_id",
				Columns:    []*schema.Column{ProductUsersColumns[0]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "product_users_user_id",
				Columns:    []*schema.Column{ProductUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ProductCategoriesColumns holds the columns for the "product_categories" table.
	ProductCategoriesColumns = []*schema.Column{
		{Name: "product_id", Type: field.TypeInt},
		{Name: "category_id", Type: field.TypeInt},
	}
	// ProductCategoriesTable holds the schema information for the "product_categories" table.
	ProductCategoriesTable = &schema.Table{
		Name:       "product_categories",
		Columns:    ProductCategoriesColumns,
		PrimaryKey: []*schema.Column{ProductCategoriesColumns[0], ProductCategoriesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "product_categories_product_id",
				Columns:    []*schema.Column{ProductCategoriesColumns[0]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "product_categories_category_id",
				Columns:    []*schema.Column{ProductCategoriesColumns[1]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CategoriesTable,
		ProductsTable,
		UsersTable,
		ProductUsersTable,
		ProductCategoriesTable,
	}
)

func init() {
	ProductUsersTable.ForeignKeys[0].RefTable = ProductsTable
	ProductUsersTable.ForeignKeys[1].RefTable = UsersTable
	ProductCategoriesTable.ForeignKeys[0].RefTable = ProductsTable
	ProductCategoriesTable.ForeignKeys[1].RefTable = CategoriesTable
}
