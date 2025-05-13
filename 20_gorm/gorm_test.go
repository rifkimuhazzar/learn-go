package gorm

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func openConnection() *gorm.DB {
	dsn := "root:n3txt.vml1@tcp(127.0.0.1:3306)/learn_go_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	return db
}

var db = openConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
}

func TestExecuteSQL(t *testing.T) {
	err := db.Exec("INSERT INTO sample(id, name) VALUES(?, ?)", "1", "React").Error
	assert.Nil(t, err)

	err = db.Exec("INSERT INTO sample(id, name) VALUES(?, ?)", "2", "Angular").Error
	assert.Nil(t, err)

	err = db.Exec("INSERT INTO sample(id, name) VALUES(?, ?)", "3", "Vue").Error
	assert.Nil(t, err)

	err = db.Exec("INSERT INTO sample(id, name) VALUES(?, ?)", "4", "Svelte").Error
	assert.Nil(t, err)
}

type Sample struct {
	ID   string
	Name string
}

func TestQuerySQL(t *testing.T) {
	var sample Sample
	err := db.Raw("SELECT id, name FROM sample WHERE id = ?", "1").Scan(&sample).Error
	assert.Nil(t, err)
	assert.Equal(t, "React", sample.Name)
	fmt.Println(sample)

	var samples []Sample
	err = db.Raw("SELECT id, name FROM sample").Scan(&samples).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(samples))
	fmt.Println(samples)
}

func TestSQLRows(t *testing.T) {
	rows, err := db.Raw("SELECT id, name FROM sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var samples []Sample
	for rows.Next() {
		var id string
		var name string

		err := rows.Scan(&id, &name)
		assert.Nil(t, err)

		samples = append(samples, Sample{
			ID:   id,
			Name: name,
		})

		fmt.Println(samples)
	}

	assert.Equal(t, 4, len(samples))
}

func TestScanRows(t *testing.T) {
	rows, err := db.Raw("SELECT id, name FROM sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var samples []Sample
	for rows.Next() {
		err := db.ScanRows(rows, &samples)
		assert.Nil(t, err)

		fmt.Println(samples)
	}

	assert.Equal(t, 4, len(samples))
}

func TestCreateUser(t *testing.T) {
	user := &User{
		ID:       "1",
		Password: "12345678",
		Name: Name{
			FirstName:  "Monkey",
			MiddleName: "D.",
			LastName:   "Luffy",
		},
		Information: "This will be ignored",
	}

	response := db.Create(user)
	assert.Nil(t, response.Error)
	assert.Equal(t, 1, int(response.RowsAffected))
}

func TestBatchInsert(t *testing.T) {
	var users []User
	for i := 2; i < 10; i++ {
		users = append(users, User{
			ID:       strconv.Itoa(i),
			Password: "12345678",
			Name:     Name{FirstName: "User-" + strconv.Itoa(i)},
		})
	}

	result := db.Create(users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 8, int(result.RowsAffected))
}

func TestTransactionSuccess(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{ID: "10", Password: "12345678", Name: Name{FirstName: "User-10"}}).Error
		if err != nil {
			return err
		}
		err = tx.Create(&User{ID: "11", Password: "12345678", Name: Name{FirstName: "User-11"}}).Error
		if err != nil {
			return err
		}
		err = tx.Create(&User{ID: "12", Password: "12345678", Name: Name{FirstName: "User-12"}}).Error
		if err != nil {
			return err
		}

		return nil
	})

	assert.Nil(t, err)
}

func TestTransactionFail(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{ID: "13", Password: "12345678", Name: Name{FirstName: "User-13"}}).Error
		if err != nil {
			return err
		}
		err = tx.Create(&User{ID: "11", Password: "12345678", Name: Name{FirstName: "User-11"}}).Error
		if err != nil {
			return err
		}

		return nil
	})

	assert.NotNil(t, err)
}

func TestManualTransactionSuccess(t *testing.T) {
	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Create(&User{ID: "13", Password: "12345678", Name: Name{FirstName: "User-13"}}).Error
	assert.Nil(t, err)

	err = tx.Create(&User{ID: "14", Password: "12345678", Name: Name{FirstName: "User-14"}}).Error
	assert.Nil(t, err)

	if err == nil {
		tx.Commit()
	}
}

func TestManualTransactionFail(t *testing.T) {
	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Create(&User{ID: "15", Password: "12345678", Name: Name{FirstName: "User-15"}}).Error
	assert.Nil(t, err)

	err = tx.Create(&User{ID: "13", Password: "12345678", Name: Name{FirstName: "User-13"}}).Error
	assert.NotNil(t, err)

	if err == nil {
		tx.Commit()
	}
}

func TestQuerySingleObject(t *testing.T) {
	user := User{}
	err := db.First(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, "1", user.ID)
	fmt.Println(user)

	user = User{}
	err = db.Last(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, "9", user.ID)
	fmt.Println(user)

	user = User{}
	err = db.Take(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, "1", user.ID)
	fmt.Println(user)
}

func TestQuerySingleObjectWithConds(t *testing.T) {
	user := User{}
	err := db.First(&user, "id = ?", "5").Error
	assert.Nil(t, err)
	assert.Equal(t, "5", user.ID)
	fmt.Println(user)

	user = User{}
	err = db.Last(&user, "id = ?", "5").Error
	assert.Nil(t, err)
	assert.Equal(t, "5", user.ID)
	fmt.Println(user)

	user = User{}
	err = db.Take(&user, "id = ?", "5").Error
	assert.Nil(t, err)
	assert.Equal(t, "5", user.ID)
	fmt.Println(user)
}

func TestQueryAllObjects(t *testing.T) {
	var users []User
	results := db.Find(&users, "id in ?", []string{"1", "2", "11", "12"})
	assert.Nil(t, results.Error)
	assert.Equal(t, 4, len(users))
	fmt.Println(users)
}

func TestQueryCondition(t *testing.T) {
	var users []User
	// default operator AND
	result := db.Where("first_name like ?", "%User%").
		Where("password = ?", "12345678").
		Find(&users)

	assert.Nil(t, result.Error)
	assert.Equal(t, 13, len(users))
}

func TestQueryConditionWithOr(t *testing.T) {
	var users []User
	result := db.Where("first_name like ?", "%User%").
		Or("password = ?", "12345678").
		Find(&users)

	assert.Nil(t, result.Error)
	assert.Equal(t, 14, len(users))
}

func TestQueryConditionWithNot(t *testing.T) {
	var users []User
	result := db.Not("first_name like ?", "%User%").
		Where("password = ?", "12345678").
		Find(&users)

	assert.Nil(t, result.Error)
	assert.Equal(t, 1, len(users))
}

func TestSelectFields(t *testing.T) {
	var users []User
	result := db.Select("id", "first_name").Find(&users)
	assert.Nil(t, result.Error)

	for _, user := range users {
		assert.NotNil(t, user.ID)
		assert.NotEqual(t, "", user.Name.FirstName)
	}

	assert.Equal(t, 14, len(users))
}

func TestStructCondition(t *testing.T) {
	userCondition := User{
		Name: Name{
			FirstName: "User-5",
			// LastName:  "", // tidak error dan tidak menjadi kondisi karena dianggap default value
		},
		Password: "12345678",
	}

	var users []User
	result := db.Where(userCondition).Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 1, len(users))
	fmt.Println(users)
}

func TestMapCondition(t *testing.T) {
	userCondition := map[string]string{
		"first_name": "User-10",
		"last_name":  "",
	}

	var users []User
	result := db.Where(userCondition).Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 1, len(users))
	fmt.Println(users)
}

func TestOrderOffsetLimit(t *testing.T) {
	var users []User
	result := db.Order("id asc, first_name asc").Offset(5).Limit(5).Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, "14", users[0].ID)
	assert.Equal(t, 5, len(users))
}

type UserResponse struct {
	// tidak perlu tag karena sudah sesuai gorm convention
	ID        string
	FirstName string
	LastName  string
}

func TestQueryNonModel(t *testing.T) {
	var users []UserResponse
	result := db.Model(&User{}).Select("id", "first_name", "last_name").Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 14, len(users))
	fmt.Println(users)
}

func TestUpdateAll(t *testing.T) {
	var user User
	result := db.Take(&user, "id = ?", "1")
	assert.Nil(t, result.Error)
	fmt.Println(user)

	user.Name.FirstName = "Roronoa"
	user.Name.MiddleName = ""
	user.Name.LastName = "Zoro"
	// user.Password = "password123"

	result = db.Save(&user)
	assert.Nil(t, result.Error)
	fmt.Println(user)
}

func TestUpdateSelectedColumns(t *testing.T) {
	result := db.Model(User{}).Where("id = ?", "1").Updates(map[string]any{
		"first_name":  "React",
		"middle_name": "",
		"last_name":   "",
	})
	assert.Nil(t, result.Error)

	result = db.Model(&User{}).Where("id = ?", "1").Update("password", "passwordDiubah")
	assert.Nil(t, result.Error)

	result = db.Where("id = ?", "1").Updates(User{
		Name: Name{
			FirstName: "Svelte",
			LastName:  "Kit",
		},
	})
	assert.Nil(t, result.Error)
}

func TestAutoIncrement(t *testing.T) {
	for range 10 {
		userLog := UserLog{
			UserID: "1",
			Action: "Test Action",
		}

		result := db.Create(&userLog)
		assert.Nil(t, result.Error)
		assert.NotEqual(t, 0, userLog.ID)
		fmt.Println(userLog)
	}
}

func TestSaveInsertOrUpdate(t *testing.T) {
	userLog := UserLog{
		UserID: "1",
		Action: "Test Action",
	}

	result := db.Save(&userLog) // insert
	assert.Nil(t, result.Error)
	fmt.Println(userLog)

	userLog.ID = 90
	result = db.Save(&userLog) // update, jika tidak ada id yang sesuai maka diinsert
	assert.Nil(t, result.Error)
	fmt.Println(userLog)
}

func TestSaveInsertOrUpdateNonAutoIncrement(t *testing.T) {
	user := User{
		ID: "90",
		Name: Name{
			FirstName: "User-90",
		},
	}

	result := db.Save(&user) // update first then if not exist do insert
	assert.Nil(t, result.Error)
	fmt.Println(user)

	user.Name.FirstName = "User-90 Updated"
	result = db.Save(&user) // update
	assert.Nil(t, result.Error)
	fmt.Println(user)
}

func TestConflict(t *testing.T) {
	user := User{
		ID: "92",
		Name: Name{
			FirstName: "User-92",
		},
	}

	result := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&user)
	assert.Nil(t, result.Error)
	fmt.Println(user)
}

func TestDelete(t *testing.T) {
	var user User

	result := db.Take(&user, "id = ?", "92")
	assert.Nil(t, result.Error)
	fmt.Println(user)

	result = db.Delete(&user)
	assert.Nil(t, result.Error)
	fmt.Println(user)

	result = db.Delete(User{}, "id = ?", "91")
	assert.Nil(t, result.Error)

	result = db.Where("id = ?", "90").Delete(User{})
	assert.Nil(t, result.Error)
}

func TestSoftDelete(t *testing.T) {
	todo := Todo{
		UserID:      "1",
		Title:       "Todo 1",
		Description: "Description 1",
	}
	fmt.Println(todo)

	result := db.Create(&todo)
	assert.Nil(t, result.Error)
	fmt.Println(todo)

	result = db.Delete(&todo)
	assert.Nil(t, result.Error)
	assert.NotNil(t, todo.DeletedAt)
	fmt.Println(todo)

	var todos []Todo
	result = db.Find(&todos)
	assert.Nil(t, result.Error)
	assert.Equal(t, 0, len(todos))
	fmt.Println(todos)
}

func TestUnscoped(t *testing.T) {
	var todo Todo

	result := db.Unscoped().Take(&todo, "id = ?", 10)
	assert.Nil(t, result.Error)
	fmt.Println(todo)

	result = db.Unscoped().Delete(&todo)
	assert.Nil(t, result.Error)
	fmt.Println(todo)

	var todos []Todo
	result = db.Unscoped().Find(&todos)
	assert.Nil(t, result.Error)
	assert.Equal(t, 8, len(todos))
	fmt.Println(todos)
}

func TestLock(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		user := User{}
		result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&user, "id = ?", "1")
		if result.Error != nil {
			return result.Error
		}

		user.Name.FirstName = "Monkey"
		user.Name.MiddleName = "D."
		user.Name.LastName = "Luffy"
		result = tx.Save(&user)
		return result.Error
	})
	assert.Nil(t, err)
}

func TestCreateWallet(t *testing.T) {
	wallet := Wallet{
		ID:      "1",
		UserID:  "1",
		Balance: 1000000,
	}
	result := db.Create(&wallet)
	assert.Nil(t, result.Error)
}

func TestRetrieveRelationWithPreload(t *testing.T) {
	user := User{}
	result := db.Preload("Wallet").Take(&user, "id = ?", "1")
	assert.Nil(t, result.Error)
	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "1", user.Wallet.ID)
	fmt.Println(user)
}

func TestRetrieveRelationWithJoins(t *testing.T) {
	user := User{}
	result := db.Joins("Wallet").Take(&user, "users.id = ?", "1")
	assert.Nil(t, result.Error)
	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "1", user.Wallet.ID)
	fmt.Println(user)
}

func TestAutoUpsert(t *testing.T) {
	user := User{
		ID:       "20",
		Password: "12345678",
		Name: Name{
			FirstName: "User-20",
		},
		Wallet: Wallet{
			ID:      "20",
			UserID:  "20",
			Balance: 2000000,
		},
	}
	result := db.Create(&user)
	assert.Nil(t, result.Error)
	assert.Equal(t, "20", user.ID)
	assert.Equal(t, "20", user.Wallet.ID)
	fmt.Println(user)
}

func TestSkipUpsert(t *testing.T) {
	user := User{
		ID:       "22",
		Password: "12345678",
		Name: Name{
			FirstName: "User-22",
		},
		Wallet: Wallet{
			ID:      "22",
			UserID:  "22",
			Balance: 2000000,
		},
	}
	result := db.Omit(clause.Associations).Create(&user)
	assert.Nil(t, result.Error)
	assert.Equal(t, "22", user.ID)
	assert.Equal(t, "22", user.Wallet.ID)
	fmt.Println(user)
}

func TestUserAndAddresses(t *testing.T) {
	user := User{
		ID:       "40",
		Password: "12345678",
		Name: Name{
			FirstName: "User-40",
		},
		Wallet: Wallet{
			ID:      "40",
			UserID:  "40",
			Balance: 3000000,
		},
		Addresses: []Address{
			{UserID: "40", Address: "Address 1"},
			{UserID: "40", Address: "Address 2"},
			{UserID: "40", Address: "Address 3"},
		},
	}

	result := db.Create(&user)
	assert.Nil(t, result.Error)
	fmt.Println(user)
}

func TestPreloadAndJoinsOneToMany(t *testing.T) {
	var usersPrealoadAndJoins []User
	result := db.Preload("Addresses").Joins("Wallet").Take(&usersPrealoadAndJoins, "users.id = ?", "30")
	// result := db.Preload("Addresses").Joins("Wallet").Take(&usersPrealoadAndJoins)
	// result := db.Joins("Addresses").Joins("Wallet").Take(&usersPrealoadAndJoins)
	// result := db.Joins("Wallet").Joins("Addresses").Find(&usersPrealoadAndJoins, "users.id = ?", "3")
	assert.Nil(t, result.Error)
	fmt.Println(usersPrealoadAndJoins)
}

func TestBelongsToManyToOne(t *testing.T) {
	fmt.Println("\nPreload")
	addresses := []Address{}
	result := db.Preload("User").Find(&addresses)
	assert.Nil(t, result.Error)
	// assert.Equal(t, 6, len(addresses))
	fmt.Println(addresses)

	fmt.Println("\nJoins")
	addresses = []Address{}
	result = db.Joins("User").Find(&addresses)
	assert.Nil(t, result.Error)
	// assert.Equal(t, 6, len(addresses))
	fmt.Println(addresses)
}

func TestBelongsToOneToOne(t *testing.T) {
	fmt.Println("\nPreload")
	wallets := []Wallet{}
	result := db.Preload("User").Find(&wallets)
	assert.Nil(t, result.Error)
	assert.Equal(t, 4, len(wallets))
	fmt.Println(wallets)

	fmt.Println("\nJoins")
	wallets = []Wallet{}
	result = db.Joins("User").Find(&wallets)
	assert.Nil(t, result.Error)
	assert.Equal(t, 4, len(wallets))
	fmt.Println(wallets)
}

func TestCreateManyToMany(t *testing.T) {
	product := Product{
		ID:    "P0001",
		Name:  "Macbook Air M4 15 16/512",
		Price: 23000000,
	}
	result := db.Create(&product)
	assert.Nil(t, result.Error)

	result = db.Table("user_like_product").Create(map[string]any{
		"user_id":    "1",
		"product_id": "P0001",
	})
	assert.Nil(t, result.Error)

	result = db.Table("user_like_product").Create(map[string]any{
		"user_id":    "2",
		"product_id": "P0001",
	})
	assert.Nil(t, result.Error)

	fmt.Println(product)
}

func TestPreloadManyToManyFromProduct(t *testing.T) {
	product := Product{}
	result := db.Preload("LikedByUsers").Take(&product, "id = ?", "P0001")
	assert.Nil(t, result.Error)
	assert.Equal(t, 2, len(product.LikedByUsers))
	fmt.Println(product)
}

func TestPreloadManyToManyFromUser(t *testing.T) {
	user := User{}
	result := db.Preload("LikedProducts").Take(&user, "id = ?", "1")
	assert.Nil(t, result.Error)
	assert.Equal(t, 1, len(user.LikedProducts))
	fmt.Println(user)
}

func TestAssociationFind(t *testing.T) {
	product := Product{}
	result := db.Take(&product, "id = ?", "P0001")
	assert.Nil(t, result.Error)
	fmt.Println(product)

	users := []User{}
	err := db.Model(&product).
		Where("users.first_name LIKE ?", "User%").
		Association("LikedByUsers").
		Find(&users)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
	fmt.Println(users)
}

func TestAssociationAppend(t *testing.T) {
	user := User{}
	result := db.Take(&user, "id = ?", "3")
	assert.Nil(t, result.Error)

	product := Product{}
	result = db.Take(&product, "id = ?", "P0001")
	assert.Nil(t, result.Error)

	err := db.Model(&user).Association("LikedProducts").Append(&product)
	assert.Nil(t, err)
}

func TestAssociationReplace(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		user := User{}
		result := tx.Take(&user, "id = ?", "1")
		if result.Error != nil {
			return result.Error
		}
		fmt.Println(user)

		wallet := Wallet{
			ID:      "100",
			UserID:  user.ID,
			Balance: 5000000,
		}
		fmt.Println(wallet)

		err := tx.Model(&user).Association("Wallet").Replace(&wallet)
		fmt.Println(wallet)
		return err
	})

	assert.Nil(t, err)
}

func TestAssociationDelete(t *testing.T) {
	user := User{}
	result := db.Take(&user, "id = ?", "3")
	assert.Nil(t, result.Error)
	fmt.Println(user)

	product := Product{}
	result = db.Take(&product, "id = ?", "P0001")
	assert.Nil(t, result.Error)
	fmt.Println(product)

	err := db.Model(&product).Association("LikedByUsers").Delete(&user)
	assert.Nil(t, err)

	fmt.Println(user)
	fmt.Println(product)
}

func TestAssociationClear(t *testing.T) {
	product := Product{}
	result := db.Take(&product, "id = ?", "P0001")
	assert.Nil(t, result.Error)

	err := db.Model(&product).Association("LikedByUsers").Clear()
	assert.Nil(t, err)
}

func TestPreloadingWithConds(t *testing.T) {
	user := User{}
	// result := db.Preload("Wallet", "balance > ?", 2000000).Find(&user)
	result := db.Preload("Wallet", "balance >= ?", 1000000).Take(&user, "id = ?", 1)
	assert.Nil(t, result.Error)
	fmt.Println(user)
}

func TestNestedPreloading(t *testing.T) {
	wallet := Wallet{}
	result := db.Preload("User.Addresses").Find(&wallet, "id = ?", 30)
	assert.Nil(t, result.Error)

	fmt.Println(wallet)
	fmt.Println(wallet.User)
	fmt.Println(wallet.User.Addresses)
}

func TestPreloadingAll(t *testing.T) {
	user := User{}
	result := db.Preload(clause.Associations).Take(&user, "id = ?", 30)
	assert.Nil(t, result.Error)
	fmt.Println(user)
}

func TestJoinQuery(t *testing.T) {
	users := []User{}
	result := db.Joins("JOIN wallets ON wallets.user_id = users.id").Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 4, len(users))
	fmt.Println(users)

	users = []User{}
	result = db.Joins("Wallet").Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 19, len(users))
	fmt.Println(users)
}

func TestJoinQueryWithConds(t *testing.T) {
	users := []User{}
	result := db.Joins("JOIN wallets ON wallets.user_id = users.id AND wallets.balance >= ?", 3000000).Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 2, len(users))
	fmt.Println(users)

	users = []User{}
	result = db.Joins("Wallet").Where("Wallet.balance >= ?", 3000000).Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 2, len(users))
	fmt.Println(users)
}

func TestCount(t *testing.T) {
	var count int64
	result := db.Model(User{}).Joins("Wallet").Where("Wallet.balance >= ?", 3000000).Count(&count)
	assert.Nil(t, result.Error)
	assert.Equal(t, int64(2), count)
	fmt.Println(count)
}

type AggregationResult struct {
	TotalBalance int64
	MinBalance   int64
	MaxBalance   int64
	AvgBalance   float64
}

func TestAggregation(t *testing.T) {
	agResult := AggregationResult{}
	dbResult := db.Model(Wallet{}).Select("sum(balance) as total_balance", "min(balance) as min_balance", "max(balance) as max_balance", "avg(balance) as avg_balance").Find(&agResult)

	assert.Nil(t, dbResult.Error)
	assert.Equal(t, int64(9000000), agResult.TotalBalance)
	assert.Equal(t, int64(1000000), agResult.MinBalance)
	assert.Equal(t, int64(3000000), agResult.MaxBalance)
	assert.Equal(t, float64(2250000), agResult.AvgBalance)

	fmt.Println(agResult)
}

func TestAggregationWithGroupAndHaving(t *testing.T) {
	agResults := []AggregationResult{}
	dbResult := db.Model(Wallet{}).Select("sum(balance) as total_balance", "min(balance) as min_balance", "max(balance) as max_balance", "avg(balance) as avg_balance").
		Joins("User").
		Group("User.id").
		Having("sum(balance) >= ?", 3000000).
		Find(&agResults)

	assert.Nil(t, dbResult.Error)
	assert.Equal(t, 2, len(agResults))

	fmt.Println(agResults)
}

func TestContext(t *testing.T) {
	ctx := context.Background()

	var user []User
	result := db.WithContext(ctx).Find(&user)
	assert.Nil(t, result.Error)
	assert.Equal(t, 19, len(user))
}

func poorWalletBalance(db *gorm.DB) *gorm.DB {
	return db.Where("balance < ?", 3000000)
}

func richWalletBalance(db *gorm.DB) *gorm.DB {
	return db.Where("balance >= ?", 3000000)
}

func TestScopes(t *testing.T) {
	var wallets []Wallet
	result := db.Scopes(poorWalletBalance).Find(&wallets)
	assert.Nil(t, result.Error)
	fmt.Println(wallets)

	result = db.Scopes(richWalletBalance).Find(&wallets)
	assert.Nil(t, result.Error)
	fmt.Println(wallets)
}

func TestMigrator(t *testing.T) {
	err := db.Migrator().AutoMigrate(&GuestBook{})
	assert.Nil(t, err)
}

func TestHook(t *testing.T) {
	user := User{
		Password: "12345678",
		Name: Name{
			FirstName: "User 100",
		},
	}

	result := db.Create(&user)
	assert.Nil(t, result.Error)
	assert.NotEqual(t, "", user.ID)
	fmt.Println(user)
}
