package viper

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	config := viper.New()
	assert.NotNil(t, config)
}

func TestConfigJSON(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "learn-go-viper", config.GetString("app.name"))
	assert.Equal(t, false, config.GetBool("app.name"))

	assert.Equal(t, true, config.GetBool("database.show_sql"))
	assert.Equal(t, "true", config.GetString("database.show_sql"))

	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.Equal(t, true, config.GetBool("database.port"))
}

func TestConfigYAML(t *testing.T) {
	config := viper.New()
	/*
		RULES
		- Dapat set config seperti di bawah:
			config.SetConfigName("config")
			config.SetConfigType("yaml")
			config.AddConfigPath(".")
		- Atau langsung dengan SetConfigFile("config.yaml") atau SetConfigFile("./config.yaml")
		- Nama file (tidak termasuk extensi) di SetConfigName() dan SetConfigFile() case-insensitive
		- Nama field saat GetString(), GetBool(), GetInt(), dll case-insensitive
		- File YAML memiliki 2 nama extensi yaitu yaml dan yml
	*/

	config.SetConfigFile("config.yaml")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "learn-go-viper", config.GetString("app.name"))
	assert.Equal(t, false, config.GetBool("app.name"))

	assert.Equal(t, true, config.GetBool("database.show_sql"))
	assert.Equal(t, "true", config.GetString("database.show_sql"))

	assert.Equal(t, 3306, config.GetInt("database.port"))
	assert.Equal(t, true, config.GetBool("database.port"))
}

func TestConfigENVFile(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.env")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "learn-go-viper", config.GetString("APP_NAME"))
	assert.Equal(t, false, config.GetBool("APP_NAME"))

	assert.Equal(t, true, config.GetBool("DATABASE_SHOW_SQL"))
	assert.Equal(t, "true", config.GetString("DATABASE_SHOW_SQL"))

	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
	assert.Equal(t, false, config.GetBool("DATABASE_PORT"))
}

func TestConfigENV(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.env")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "learn-go-viper", config.GetString("APP_NAME"))
	assert.Equal(t, false, config.GetBool("APP_NAME"))

	assert.Equal(t, true, config.GetBool("DATABASE_SHOW_SQL"))
	assert.Equal(t, "true", config.GetString("DATABASE_SHOW_SQL"))

	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
	assert.Equal(t, false, config.GetBool("DATABASE_PORT"))

	/*
		- Jika menggunakan AutomaticEnv() maka akan mengutamakan mengambil env vars dari shell/os
		- Jika tidak ada dari shell/os maka akan mengambil dari file configuration
	*/
	config.AutomaticEnv()
	// assert.Equal(t, "Hello", config.GetString("FROM_ENV"))
	// assert.Equal(t, 3000, config.GetString("DATABASE_PORT"))
}

/*
	JENIS-JENIS ENV VARS DI WINDOWS, MACOS, DAN LINUX
	- system/global env, user-specific env, dan process env
	- prioritas
		- proccess (temporary -> permanent)
		- user-specific
		- system/global
	- env vars shell permanent dan temporary (process env)
		- permanent ditulis di file configuration masing-masing shell
		- temporary ditulis langsung di dalam terminal tanpa disimpan ke file configurationnya
	- env vars di project aplikasi/web (process/project env)
	- env vars untuk project app/web dapat di export/source ke proccess env temporary
	- urutan prioritas dapat bervariasi tergantung bagaimana app/web atau shell membaca dan memuat variabelnya, misal saat menggunakan method AutomaticEnv() maka akan memprioritaskan dari shell/os
*/

/*
	- Windows env vars
		- user variables (user-specific)
		- system variables (global)
		- shell specific dan user specific,
			tergantung nama file konfigurasi dan lokasinya berdasarkan shell masing-masing,
			ini disebut proccess env (dapat permanent atau temporary)
	- MacOS/Linux env vars
		- user (file konfigurasi shell seperti ~/.bashrc, ~/.bash_profile, dll)
		- global (file konfigurasi shell seperti /etc/profile, /etc/environment, dll)
		- env vars user/global dapat berlaku untuk satu shell atau beberapa shell tergantung jenis shell
		- konfigurasi di /etc/environment dapat berlaku untuk semua shell di semua user
*/
