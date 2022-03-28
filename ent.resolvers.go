package bitsports

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bitsports/driver-framework/middlewares"
	"bitsports/ent"
	"bitsports/ent/user"
	contexthandler "bitsports/usecase/context_handler"
	"bitsports/usecase/jwt"
	"context"
	"errors"
	"log"
	"net/http"
)

func (r *mutationResolver) CreateUser(ctx context.Context, user UserInputSingUp) (*ent.User, error) {

	if err := r.passwordValidator.ValidatePassword([]byte(user.Password)); err != nil {
		return nil, err
	}

	cryptedPassword, err := r.passwordValidator.EncryptPassword([]byte(user.Password))
	if err != nil {
		return nil, err
	}

	return r.client.User.Create().
		SetEmail(user.Email).
		SetName(user.Name).
		SetPassword(string(cryptedPassword)).
		Save(ctx)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, user UserInputUpdate) (*ent.User, error) {
	userClaims, ok := ctx.Value(contexthandler.UserContextKey).(*ent.User)
	if !ok {
		return nil, errors.New("cant get user from ctx")
	}

	uploader := r.client.User.UpdateOneID(userClaims.ID)
	if user.Email != nil {
		uploader.SetEmail(*user.Email)
	}
	if user.Name != nil {
		uploader.SetName(*user.Name)
	}
	if user.Password != nil {
		if err := r.passwordValidator.ValidatePassword([]byte(*user.Password)); err != nil {
			return nil, err
		}
		cryptedPassword, err := r.passwordValidator.EncryptPassword(
			[]byte(*user.Password),
		)
		if err != nil {
			return nil, err
		}
		uploader.SetPassword(string(cryptedPassword))
	}

	userUploaded, err := uploader.Save(ctx)
	if err != nil {
		return nil, err
	}
	return userUploaded, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*ent.User, error) {
	user, err := r.client.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := r.client.User.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *mutationResolver) AddProducts(ctx context.Context, productIds []int) (*ent.User, error) {
	user, err := contexthandler.GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return user.Update().AddProductIDs(productIds...).Save(ctx)
}

func (r *mutationResolver) CreateProduct(ctx context.Context, product ProductInput) (*ent.Product, error) {
	return r.client.Product.Create().
		SetName(product.Name).
		SetPrice(product.Price).
		Save(ctx)
}

func (r *mutationResolver) UpdateProduct(ctx context.Context, productInput ProductInputUpdate) (*ent.Product, error) {
	productSaved, err := r.client.Product.Get(ctx, productInput.ID)
	if err != nil {
		return nil, err
	}

	updater := productSaved.Update()
	if productInput.Name != nil {
		updater.SetName(*productInput.Name)
	}
	if productInput.Price != nil {
		updater.SetPrice(*productInput.Price)
	}
	productSaved, err = updater.Save(ctx)
	if err != nil {
		return nil, err
	}
	return productSaved, nil
}

func (r *mutationResolver) DeleteProduct(ctx context.Context, id int) (*ent.Product, error) {
	product, err := r.client.Product.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	err = r.client.Product.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *mutationResolver) AddCategories(ctx context.Context, id int, categoryIds []int) (*ent.Product, error) {
	product, err := r.client.Product.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return product.Update().AddCategoryIDs(categoryIds...).Save(ctx)
}

func (r *mutationResolver) CreateCategory(ctx context.Context, category CategoryInput) (*ent.Category, error) {
	return r.client.Category.Create().
		SetName(category.Name).
		Save(ctx)
}

func (r *mutationResolver) UpdateCategory(ctx context.Context, categoryInput CategoryInputUpdate) (*ent.Category, error) {
	categorySaved, err := r.client.Category.Get(ctx, categoryInput.ID)
	if err != nil {
		return nil, err
	}

	updater := categorySaved.Update()
	if categoryInput.Name != nil {
		updater.SetName(*categoryInput.Name)
	}
	categorySaved, err = updater.Save(ctx)
	if err != nil {
		return nil, err
	}
	return categorySaved, nil
}

func (r *mutationResolver) DeleteCategory(ctx context.Context, id int) (*ent.Category, error) {
	category, err := r.client.Category.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	err = r.client.Category.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (r *queryResolver) AllUsers(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().All(ctx)
}

func (r *queryResolver) AllProducts(ctx context.Context) ([]*ent.Product, error) {
	return r.client.Product.Query().All(ctx)
}

func (r *queryResolver) AllCategories(ctx context.Context) ([]*ent.Category, error) {
	return r.client.Category.Query().All(ctx)
}

func (r *queryResolver) Login(ctx context.Context, userSingIn UserInputSingIn) (*ent.User, error) {
	userSaved, err := r.client.User.Query().Where(
		user.Email(userSingIn.Email),
	).Only(ctx)
	if err != nil {
		return nil, err
	}

	if err := r.passwordValidator.ComparePassword(
		[]byte(userSingIn.Password),
		[]byte(userSaved.Password),
	); err != nil {
		return nil, err
	}

	userClaims := jwt.UserClaims{
		ID: userSaved.ID,
	}

	token, err := jwt.CreateToken(userClaims)
	if err != nil {
		log.Fatal(err)
	}
	responser := ctx.Value(middlewares.ResponseContextKey).(*http.ResponseWriter)
	http.SetCookie(*responser, &http.Cookie{
		Name:  "jwt",
		Value: token,
	})
	return userSaved, nil
}

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, id []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, id)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
