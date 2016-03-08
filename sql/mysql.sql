CREATE TABLE Administrator (
	id INTEGER PRIMARY KEY AUTO_INCREMENT,
	login_id VARCHAR(255) NOT NULL UNIQUE,
	password VARCHAR(255) NOT NULL
)ENGINE=InnoDB;

CREATE TABLE Organization (
	id INTEGER PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(255) NOT NULL UNIQUE
)ENGINE=InnoDB;

CREATE TABLE Field (
	id INTEGER PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(255) NOT NULL UNIQUE,
	org_id INTEGER NOT NULL
)ENGINE=InnoDB;

CREATE TABLE Analysis (
	id INTEGER PRIMARY KEY AUTO_INCREMENT,
	field_id INTEGER NOT NULL,
	field_key VARCHAR(255) UNIQUE,
	analysis_date DATE NOT NULL,
	ph DOUBLE(2,1),
	phk DOUBLE(2,1),
	ec DOUBLE(3,2),
	php DOUBLE(3,1),		#リン酸吸収係数
	eofph DOUBLE(3,1),		#有効態リン酸
	k DOUBLE(3,1),			#カリ
	dk DOUBLE(3,1),			#カリ飽和度
	ca DOUBLE(3,1),			#石灰
	dca DOUBLE(3,1),		#石灰飽和度
	mg DOUBLE(3,1),			#苦土
	dmg DOUBLE(3,1),		#苦土飽和度
	cec INTEGER,
	dcec DOUBLE(3,1),		#塩基飽和度
	capermg DOUBLE(2,2),	#Ca/Mg
	mgperk DOUBLE(2,2)		#Mg/k
)ENGINE=InnoDB;