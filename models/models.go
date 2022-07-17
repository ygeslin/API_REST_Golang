package models
type Friend struct{
    Id          int        `json:"id"`
    Name        string     `json:"name"`
}
type User struct {
    // _id         string        `json:"_id,omitempty"`
    ID          string        `json:"id"`
    Password    string        `json:"password"`
    IsActive    bool          `json:"isActive"`
    Balance     string        `json:"balance"`
    Age         interface{}   `json:"age"`
    Name        string        `json:"name"`
    Gender      string        `json:"gender"`
    Company     string        `json:"company"`
    Email       string        `json:"email"`
    Phone       string        `json:"phone"`
    Address     string        `json:"address"`
    About       string        `json:"about"`
    Registered  string        `json:"registered"`
    Latitude    float32       `json:"Latitude"`
    Longitude   float32       `json:"Longitude"`
    Tags        []string      `json:"tags"`
    Friends     []Friend      `json:"friends"`
    Data        string        `json:"data"`
}
