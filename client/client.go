package client

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	DatabaseName = ""
	DatabaseURI  = ""
)

type KatClient struct {
	db           *mongo.Client
	databaseName string
	kctx         context.Context
}

type AccessoryData struct {
	Placement string           `bson:"placement" json:"placement"`
	Accessory map[int][]string `bson:"accessory" json:"accessory"`
	SubType   string           `bson:"subtype" json:"subType"`
}

type AccessoryDataRequest struct {
	Data AccessoryData `json:"data"`
}

type PaletteData struct {
	Palette []string `bson:"palette" json:"palette"`
	Type    string   `bson:"type" json:"type"`
}

type AccessoryResponse struct {
	Accessory *AccessoryData `json:"accessory"`
	Palette   *PaletteData   `json:"palette"`
}

type FaceData struct {
	Face map[int][]string `bson:"face"`
}

type FaceResponse struct {
	Face    *FaceData    `json:"face"`
	Palette *PaletteData `json:"palette"`
}

type BackgroundData struct {
	Background map[int][]string `bson:"background"`
}

type BackgroundResponse struct {
	Background *BackgroundData `json:"background"`
}

type KatAccessories struct {
	Top    *AccessoryResponse `json:"top"`
	Mid    *AccessoryResponse `json:"mid"`
	Bottom *AccessoryResponse `json:"bottom"`
}

type KatResponse struct {
	Face        *FaceResponse       `json:"face"`
	Background  *BackgroundResponse `json:"background"`
	Accessories *KatAccessories     `json:"accessories"`
}

func NewClient() (*KatClient, error) {
	dbClient, err := mongo.NewClient(options.Client().ApplyURI(DatabaseURI))
	if err != nil {
		return nil, errors.New("unable to create db client")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()
	if err := dbClient.Connect(ctx); err != nil {
		panic("unable to connect to db")
	} else {
		fmt.Println("connected to db")
	}
	return &KatClient{
		db:           dbClient,
		databaseName: "cryptokat",
		kctx:         ctx,
	}, nil
}

func GenerateRandomAccessoryIndex(aData []AccessoryData) int {
	return rand.Intn(len(aData))
}
func GenerateRandomPaletteIndex(pData []PaletteData) int {
	return rand.Intn(len(pData))
}
func GenerateRandomFaceIndex(fData []FaceData) int {
	return rand.Intn(len(fData))
}
func GenerateRandomBackgroundIndex(bData []BackgroundData) int {
	return rand.Intn(len(bData))
}

func (k KatClient) GetAccessoryData(placement string) (*AccessoryResponse, error) {
	var accessories []AccessoryData
	fmt.Printf("placement in client %s\n", placement)
	nctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	katDB := k.db.Database(k.databaseName)
	accessoryCollection := katDB.Collection("accessory")

	cursor, err := accessoryCollection.Find(nctx, bson.D{{"placement", placement}})
	defer cursor.Close(nctx)
	if err != nil {
		log.Fatalf("cursor err %v", err)
		return nil, errors.New(fmt.Sprintf("error: %v", err))
	} else {
		for cursor.Next(nctx) {
			var result AccessoryData
			if err := cursor.Decode(&result); err != nil {
				fmt.Println("cursor.Next() error:", err)
				os.Exit(1)
			} else {
				accessories = append(accessories, result)
			}
		}
	}

	if len(accessories) == 0 {
		return nil, nil
	}
	rIndex := GenerateRandomAccessoryIndex(accessories)

	chosenAccess := &AccessoryData{
		Accessory: accessories[rIndex].Accessory,
		Placement: accessories[rIndex].Placement,
		SubType:   accessories[rIndex].SubType,
	}

	pd, err := k.GetPaletteData("accessory")
	if err != nil {
		fmt.Println("no palettes found for accessory")
	}
	out := &AccessoryResponse{
		Accessory: chosenAccess,
		Palette:   pd,
	}
	return out, nil
}

func (k KatClient) GetFaceData() (*FaceResponse, error) {
	var faces []FaceData
	nctx, cancel := context.WithCancel(k.kctx)
	defer cancel()
	katDB := k.db.Database(k.databaseName)
	faceCollection := katDB.Collection("face")
	cursor, err := faceCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		defer cursor.Close(nctx)
	} else {
		for cursor.Next(nctx) {
			var result FaceData
			if err := cursor.Decode(&result); err != nil {
				fmt.Println("cursor.Next() error:", err)
				os.Exit(1)
			} else {
				faces = append(faces, result)
			}
		}
	}

	fIndex := GenerateRandomFaceIndex(faces)
	chosenFace := &FaceData{
		Face: faces[fIndex].Face,
	}
	pd, err := k.GetPaletteData("face")
	if err != nil {
		fmt.Println("no palettes found for accessory")
	}
	out := &FaceResponse{
		Face:    chosenFace,
		Palette: pd,
	}
	return out, nil
}

func (k KatClient) GetBackgroundData() (*BackgroundResponse, error) {
	var backgrounds []BackgroundData
	nctx, cancel := context.WithCancel(k.kctx)
	defer cancel()
	katDB := k.db.Database(k.databaseName)
	backgroundCollection := katDB.Collection("background")
	cursor, err := backgroundCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		defer cursor.Close(nctx)
	} else {
		for cursor.Next(nctx) {
			var result BackgroundData
			if err := cursor.Decode(&result); err != nil {
				fmt.Println("cursor.Next() error:", err)
				os.Exit(1)
			} else {
				backgrounds = append(backgrounds, result)
			}
		}
	}
	bIndex := GenerateRandomBackgroundIndex(backgrounds)
	chosenBackground := &BackgroundData{
		Background: backgrounds[bIndex].Background,
	}

	out := &BackgroundResponse{
		Background: chosenBackground,
	}
	return out, nil
}

func (k KatClient) GetPaletteData(itemType string) (*PaletteData, error) {
	var palettes []PaletteData
	fmt.Printf("subtype palette %s\n", itemType)
	nctx, cancel := context.WithCancel(k.kctx)
	defer cancel()
	katDB := k.db.Database(k.databaseName)
	paletteCollection := katDB.Collection("palette")

	cursor, err := paletteCollection.Find(context.TODO(), bson.D{{"type", itemType}})
	if err != nil {
		defer cursor.Close(nctx)
	} else {
		for cursor.Next(nctx) {
			var result PaletteData
			if err := cursor.Decode(&result); err != nil {
				fmt.Println("cursor.Next() error:", err)
				os.Exit(1)
			} else {
				palettes = append(palettes, result)
			}
		}
	}
	if len(palettes) == 0 {
		return nil, errors.New("no palettes found for accessory/face")
	}
	pIndex := GenerateRandomPaletteIndex(palettes)

	return &PaletteData{
		Palette: palettes[pIndex].Palette,
		Type:    itemType,
	}, nil
}

func (k KatClient) CreateAccessoryData(a *AccessoryData) error {
	katDB := k.db.Database(k.databaseName)
	accessoryCollection := katDB.Collection("accessory")
	result, err := accessoryCollection.InsertOne(context.TODO(), a)
	if err != nil {
		errors.New(fmt.Sprintf("unable to insert in accessory collection, error: %v\n", err))
	}
	fmt.Printf(fmt.Sprintf("Inserted: %v into db\n", result))
	return nil
}

func (k KatClient) CreatePaletteData(p *PaletteData) error {
	fmt.Printf("incoming palette data %v\n", p)
	katDB := k.db.Database(k.databaseName)
	paletteCollection := katDB.Collection("palette")
	result, err := paletteCollection.InsertOne(context.TODO(), p)
	if err != nil {
		errors.New(fmt.Sprintf("unable to insert in palette collection error: %v\n", err))
	}
	fmt.Printf(fmt.Sprintf("Inserted: %v into db\n", result))
	return nil
}

func (k KatClient) GetKat() (*KatResponse, error) {
	at, _ := k.GetAccessoryData("top")
	am, _ := k.GetAccessoryData("mid")
	ab, _ := k.GetAccessoryData("bottom")


	ka := &KatAccessories{
		Top:   	at,
		Mid:    am,
		Bottom: ab,
	}

	f, ferr := k.GetFaceData()
	b, berr := k.GetBackgroundData()

	if ferr != nil || berr != nil {
		return nil, errors.New("unable to retrieve face or background")
	}

	out := &KatResponse{
		Accessories: ka,
		Face:        f,
		Background:  b,
	}
	return out, nil
}
