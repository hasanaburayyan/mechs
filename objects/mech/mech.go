package mech

type Armor struct {
	HitPoints int
}

type Item struct {
	Name string
	Description string
}

type Weapon struct {
	Name string
	Description string
	Location string
	Hit int
	Damage int
}

type StorageSpace struct {
	Slots int
	Items []Item
	Weapons []Weapon
}

type Piece struct {
	Armor Armor
	Storage StorageSpace
}

type MechWarrior struct {
	Head Piece
	RightArm Piece
	RightTorso Piece
	LeftArm Piece
	LeftTorso Piece
	Chest Piece
	LeftLeg Piece
	RightLeg Piece
	RearRightTorso Armor
	RearTorso Armor
	RearLeftTorso Armor
}