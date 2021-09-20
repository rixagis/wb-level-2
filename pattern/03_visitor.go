// Применение паттерна "посетитель" на примере классов географических объектов (город, озеро и гора) и разных операций с ними.
package pattern

import (
	"fmt"
)

// ========================== Посещаемые объекты =====================================

// Visitable - интерфейс объекта, который может принять посетителя
type Visitable interface {
	Accept(Visitor)
}

// City - структура города
type City struct {
	Name string
	FoundingYear int
	Country string
}

// Accept принимает посетителя
func (c City) Accept(v Visitor) {
	v.VisitCity(c)
}

// Lake - структура озера
type Lake struct {
	Name string
	MaxDepth int	// макс.глубина в метрах
	Area int		// площадь в кв.км.
}

// Accept принимает посетителя
func (l Lake) Accept(v Visitor) {
	v.VisitLake(l)
}

// Mountain - структура горы
type Mountain struct {
	Name string
	Height int		// высота в метрах
	Range string
}

// Accept принимает посетителя
func (f Mountain) Accept(v Visitor) {
	v.VisitMountain(f)
}


// ===================================== Посетители ===========================================

// Visitor - интерфейс посетителей
type Visitor interface {
	VisitCity(City)
	VisitLake(Lake)
	VisitMountain(Mountain)
}

// PrintDescriptionVisitor - посетитель, который составляет подробное описание об объекте
type PrintDescriptionVisitor struct {
	description string
}

// VisitCity посещает город
func (v *PrintDescriptionVisitor) VisitCity(c City) {
	v.description = fmt.Sprintf("The city of %s, founded in %d. Currently belongs to %s.", c.Name, c.FoundingYear, c.Country)
}

// VisitLake посещает озеро
func (v *PrintDescriptionVisitor) VisitLake(l Lake) {
	v.description = fmt.Sprintf("Lake %s. Area: %d km.sq., maxinum depth: %d m.", l.Name, l.Area, l.MaxDepth)
}

// VisitMountain посещает гору
func (v *PrintDescriptionVisitor) VisitMountain(m Mountain) {
	v.description = fmt.Sprintf("Mount %s. Height: %d m. Part of the %s mountains.", m.Name, m.Height, m.Range)
}

// GetDescription возвращает описание, составленное после посещения
func (v *PrintDescriptionVisitor) GetDescription() string {
	return v.description
}



// PrintShortInfoVisitor - посетитель, составляющи короткое описание объектов
type PrintShortInfoVisitor struct {
	info string
}

// VisitCity посещает город
func (v *PrintShortInfoVisitor) VisitCity(c City) {
	v.info = fmt.Sprintf("%s city (%s)", c.Name, c.Country)
}

// VisitLake посещает озеро
func (v *PrintShortInfoVisitor) VisitLake(l Lake) {
	v.info = fmt.Sprintf("%s lake (%d km.sq.)", l.Name, l.Area)
}

// VisitMountain посещает гору
func (v *PrintShortInfoVisitor) VisitMountain(m Mountain) {
	v.info = fmt.Sprintf("Mount %s (%d m)", m.Name, m.Height)
}

// GetInfo вовзращает описание, составленное после посещения
func (v *PrintShortInfoVisitor) GetInfo() string {
	return v.info
}



// SortByCategoryVisitor - посетитель, сортирующий объекты интерфейса Visitable по категориям: города, озера и горы
type SortByCategoryVisitor struct {
	Cities []City
	Lakes []Lake
	Mountains []Mountain
}

// VisitCity посещает город
func (v *SortByCategoryVisitor) VisitCity(c City) {
	v.Cities = append(v.Cities, c)
}

// VisitLake посещает озеро
func (v *SortByCategoryVisitor) VisitLake(l Lake) {
	v.Lakes = append(v.Lakes, l)
}

// VisitMountain посещает гору
func (v *SortByCategoryVisitor) VisitMountain(m Mountain) {
	v.Mountains = append(v.Mountains, m)
}


// Пример использования
/*func main() {
	// исходные разнородные данные в одном слайсе Visitable
	var geographicObjects = []Visitable{
		City{
			Name: "Moscow",
			Country: "Russia",
			FoundingYear: 1147,
		},
		City{
			Name: "New York",
			Country: "USA",
			FoundingYear: 1624,
		},
		Lake{
			Name: "Baikal",
			Area: 31722,
			MaxDepth: 1642,
		},
		Lake{
			Name: "Michigan",
			Area: 58030,
			MaxDepth: 281,
		},
		Mountain{
			Name: "Everest",
			Height: 8849,
			Range: "Himalayas",
		},
		Mountain{
			Name: "Elbrus",
			Height: 5642,
			Range: "Caucasus",
		},
	}

	// посетители
	var (
		descriptionVisitor = &PrintDescriptionVisitor{}
		infoVisitor = &PrintShortInfoVisitor{}
		sortVisitor = &SortByCategoryVisitor{}
	)

	// демонстрация PrintDescriptionVisitor
	fmt.Println("Long descriptions:")
	for _, obj := range geographicObjects {
		obj.Accept(descriptionVisitor)
		fmt.Println(descriptionVisitor.GetDescription())
	}

	// демонстрация PrintShortInfoVisitor
	fmt.Println()
	fmt.Println("Short info:")
	for _, obj := range geographicObjects {
		obj.Accept(infoVisitor)
		fmt.Println(infoVisitor.GetInfo())
	}

	// демонстрация SortByCategoryVisitor
	fmt.Println()
	for _, obj := range geographicObjects {
		obj.Accept(sortVisitor)
	}
	fmt.Print("Cities: ")
	for _, city := range sortVisitor.Cities {
		fmt.Print(city.Name, " ")
	}
	fmt.Println()
	fmt.Print("Lakes: ")
	for _, lake := range sortVisitor.Lakes {
		fmt.Print(lake.Name, " ")
	}
	fmt.Println()
	fmt.Print("Mountains: ")
	for _, mountain := range sortVisitor.Mountains {
		fmt.Print(mountain.Name, " ")
	}

}*/