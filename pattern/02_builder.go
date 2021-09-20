// Реализация паттерна "строитель" на примере строителя объектов JSON.
// В этом пакете реализованы типы объектов и других типов данных JSON и строители объектов и массивов JSON
// по аналогии с пакетом javax.json из стандартной библиотеки языка Java.
package pattern

import (
	"errors"
)

// Константы типов данных JSON
const (
	JsonTypeNumber = iota
	JsonTypeString
	JsonTypeBoolean
	JsonTypeNull
	JsonTypeArray
	JsonTypeObject
)

// Собственные ошибки
var (
	ErrInvalidArrayIndex = errors.New("invalid array index")		//неправильно указан индекс в массиве JsonArray
	ErrInvalidFieldName = errors.New("invalid object field name")	//не найдено имя поля в JsonObject
)

// JsonValue - интерфейс для размещения значений разных типов JSON в одной коллекции
type JsonValue interface {
	Type() int
}

// ================= JsonNumber ====================

// JsonNumber соответствует типу number JSON
type JsonNumber struct {
	value float64
}

// Type возвращает константу типа
func (number JsonNumber) Type() int {
	return JsonTypeNumber
}

// Value возвращает значение
func (number JsonNumber) Value() float64 {
	return number.value
}

// ================= JsonString =======================

// JsonString соответствует типу string JSON
type JsonString struct {
	value string
}

// Type возвращает константу типа
func (str JsonString) Type() int {
	return JsonTypeString
}

// Value возвращает значение
func (str JsonString) Value() string {
	return str.value
}

// ===================== JsonBoolean =====================

// JsonBoolean соответствует типу number JSON
type JsonBoolean struct {
	value bool
}

// Type возвращает константу типа
func (b JsonBoolean) Type() int {
	return JsonTypeBoolean
}

// Value возвращает значение
func (b JsonBoolean) Value() bool {
	return b.value
}

// ==================== JsonNull ========================

// JsonBoolean соответствует типу null JSON
type JsonNull struct {
}

// Type возвращает константу типа
func (n JsonNull) Type() int {
	return JsonTypeNull
}

// ===================== JsonArray =========================

// JsonBoolean соответствует типу array JSON
type JsonArray struct {
	elements []JsonValue
}

// Type возвращает константу типа
func (arr JsonArray) Type() int {
	return JsonTypeArray
}

// Len возвращает длину массива
func (arr JsonArray) Len() int {
	return len(arr.elements)
}

// GetElement возвращает элемент с индексом index. Ошибка ErrInvalidArrayIndex, если индекс неверен
func (arr JsonArray) GetElement(index int) (JsonValue, error) {
	if index < 0 || index > len(arr.elements) {
		return nil, ErrInvalidArrayIndex
	}
	return arr.elements[index], nil
}

// =================== JsonObject =========================

// JsonObject соответствует типу object JSON
type JsonObject struct {
	fields map[string]JsonValue
}

// Type возвращает константу типа
func (obj JsonObject) Type() int {
	return JsonTypeObject
}

// Len возвращает длину массива
func (obj JsonObject) Len() int {
	return len(obj.fields)
}

// GetField возвращает значение поля с именем name. Ошибка ErrInvalidFieldName, если поля с таким именем нет
func (obj JsonObject) GetField(name string) (JsonValue, error) {
	var res, ok = obj.fields[name]
	if !ok {
		return nil, ErrInvalidFieldName
	}
	return res, nil
}

// =================== JsonObjectBuilder ============================

// JsonObjectBuilder - строитель объектов JsonObject
type JsonObjectBuilder struct {
	fields map[string]JsonValue
}

// NewJsonObjectBuilder - конструктор строителя
func NewJsonObjectBuilder() *JsonObjectBuilder {
	return &JsonObjectBuilder{make(map[string]JsonValue)}
}

// AddInt добавляет поле типа int, будет храниться, как JsonNumber
func (builder *JsonObjectBuilder) AddInt(name string, value int) *JsonObjectBuilder {
	builder.fields[name] = JsonValue(JsonNumber{float64(value)})
	return builder
}

// AddFloat64 добавляет поле типа float64, будет храниться, как JsonNumber
func (builder *JsonObjectBuilder) AddFloat64(name string, value float64) *JsonObjectBuilder {
	builder.fields[name] = JsonValue(JsonNumber{value})
	return builder
}

// AddString добавляет поле типа string, будет храниться, как JsonString
func (builder *JsonObjectBuilder) AddString(name string, value string) *JsonObjectBuilder {
	builder.fields[name] = JsonValue(JsonString{value})
	return builder
}

// AddBool добавляет поле типа bool, будет храниться, как JsonBoolean
func (builder *JsonObjectBuilder) AddBool(name string, value bool) *JsonObjectBuilder {
	builder.fields[name] = JsonValue(JsonBoolean{value})
	return builder
}

// AddNull добавляет поле типа Null, будет храниться, как JsonNull
func (builder *JsonObjectBuilder) AddNull(name string) *JsonObjectBuilder {
	builder.fields[name] = JsonValue(JsonNull{})
	return builder
}

// AddObject добавляет поле типа JsonObject
func (builder *JsonObjectBuilder) AddObject(name string, value *JsonObject) *JsonObjectBuilder {
	builder.fields[name] = JsonValue(*value)
	return builder
}

// AddArray добавляет поле типа JsonArray
func (builder *JsonObjectBuilder) AddArray(name string, value *JsonArray) *JsonObjectBuilder {
	builder.fields[name] = JsonValue(*value)
	return builder
}

// Build создает результирующий объект JsonObject и возвращает указатель на него
func (builder *JsonObjectBuilder) Build() *JsonObject {
	var result = JsonObject{make(map[string]JsonValue)}
	for key, value := range builder.fields {
		result.fields[key] = value
	}

	return &result
}


// =================== JsonArrayBuilder ============================

// JsonArrayBuilder - строитель массивов типа JsonArray
type JsonArrayBuilder struct {
	fields []JsonValue
}

// NewJsonArrayBuilder - конструктор строителя массивов
func NewJsonArrayBuilder() *JsonArrayBuilder {
	return &JsonArrayBuilder{make([]JsonValue, 0)}
}

// AddInt добавляет элемент типа int, будет храниться как JsonNumber
func (builder *JsonArrayBuilder) AddInt(value int) *JsonArrayBuilder {
	builder.fields = append(builder.fields, JsonValue(JsonNumber{float64(value)}))
	return builder
}

// AddFloat64 добавляет элемент типа float64, будет храниться как JsonNumber
func (builder *JsonArrayBuilder) AddFloat64(value float64) *JsonArrayBuilder {
	builder.fields = append(builder.fields, JsonValue(JsonNumber{value}))
	return builder
}

// AddString добавляет элемент типа string
func (builder *JsonArrayBuilder) AddString(value string) *JsonArrayBuilder {
	builder.fields = append(builder.fields, JsonValue(JsonString{value}))
	return builder
}

// AddBool добавляет элемент типа bool
func (builder *JsonArrayBuilder) AddBool(value bool) *JsonArrayBuilder {
	builder.fields = append(builder.fields, JsonValue(JsonBoolean{value}))
	return builder
}

// AddNull добавляет элемент типа null
func (builder *JsonArrayBuilder) AddNull() *JsonArrayBuilder {
	builder.fields = append(builder.fields, JsonValue(JsonNull{}))
	return builder
}

// AddObject добавляет элемент типа JsonObject
func (builder *JsonArrayBuilder) AddObject(value *JsonObject) *JsonArrayBuilder {
	builder.fields = append(builder.fields, JsonValue(*value))
	return builder
}

// AddArray добавляет элемент типа JsonArray
func (builder *JsonArrayBuilder) AddArray(value *JsonArray) *JsonArrayBuilder {
	builder.fields = append(builder.fields, JsonValue(*value))
	return builder
}

// Build создает объект массива типа JsonArray и возвращает указатель на него
func (builder *JsonArrayBuilder) Build() *JsonArray {
	var result = JsonArray{make([]JsonValue, 0)}
	for _, value := range builder.fields {
		result.elements = append(result.elements, value)
	}

	return &result
}


// Пример использования
/*func main() {
	var builder = NewJsonObjectBuilder()
	var object = builder.AddString("firstName", "John").
						AddString("lastName", "Smith").
						AddInt("age", 25).
						AddObject("address", NewJsonObjectBuilder().
							AddString("streetAddress", "21 2nd Street").
							AddString("city", "New-York").
							AddString("state", "NY").
							AddString("postalCode", "10021").
							Build()).
						AddArray("phoneNumber", NewJsonArrayBuilder().
							AddObject(NewJsonObjectBuilder().
								AddString("type", "home").
								AddString("number", "212 555-1234").
								Build()).
							AddObject(NewJsonObjectBuilder().
								AddString("type", "fax").
								AddString("number", "646 555-4567").
								Build()).
							Build()).
						Build()
	
	fmt.Println("Built object:")
	fmt.Println(*object)
}
*/