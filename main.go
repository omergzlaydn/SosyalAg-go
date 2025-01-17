package main

import (
    "fmt"
    "math/rand"
    "time"
)

// Sosyal ağdaki kullanıcı bağlantılarını tutan bir harita. bunu global tanımlamamız tüm fonksiyonların buna erişmesini sağlayacaktır
var baglantilar = make(map[int][]int)

// BaglantiOlustur iki kullanıcı arasında karşılıklı bağlantı oluşturur
func BaglantiOlustur(a, b int) {
    baglantilar[a] = append(baglantilar[a], b)
    baglantilar[b] = append(baglantilar[b], a)
}

// BaglantiKopar iki kullanıcı arasındaki bağlantıyı koparır
func BaglantiKopar(a, b int) {
    baglantilar[a] = remove(baglantilar[a], b)
    baglantilar[b] = remove(baglantilar[b], a)
}

// remove bir slice'tan belirli bir öğeyi kaldırır
func remove(slice []int, elem int) []int {
    for i, v := range slice {
        if v == elem {
            return append(slice[:i], slice[i+1:]...)
        }
    }
    return slice
}

// YuzKullaniciOlustur 100 farklı kullanıcı oluşturur
func YuzKullaniciOlustur() {
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < 100; i++ {
        id := rand.Intn(100000000)
        // Benzersiz kullanıcı id'si ekleyene kadar döngü
        for _, exists := baglantilar[id]; exists; id = rand.Intn(100000000) {
        }
        baglantilar[id] = []int{}
    }
}

// OrtakBaglantiTespiti iki kullanıcının ortak arkadaşlarını bulur
func OrtakBaglantiTespiti(a, b int) []int {
    var ortaklar []int
    bagA := make(map[int]bool)
    for _, friend := range baglantilar[a] {
        bagA[friend] = true
    }
    for _, friend := range baglantilar[b] {
        if bagA[friend] {
            ortaklar = append(ortaklar, friend)
        }
    }
    return ortaklar
}

func main() {
    // Kullanıcıları oluştur ve birbirleriyle bağlantı kur
    BaglantiOlustur(55005500, 45001500)
    BaglantiOlustur(55005500, 35021300)
    BaglantiOlustur(45001500, 35021300)

    // Bağlantıları görüntüle
    fmt.Println("55005500 bağlantıları:", baglantilar[55005500])
    fmt.Println("45001500 bağlantıları:", baglantilar[45001500])
    fmt.Println("35021300 bağlantıları:", baglantilar[35021300])

    // 100 farklı kullanıcı oluştur
    YuzKullaniciOlustur()
    fmt.Println("100 kullanıcı oluşturuldu:", len(baglantilar))

    // Ortak bağlantı tespiti
    ortaklar := OrtakBaglantiTespiti(55005500, 45001500)
    fmt.Println("55005500 ve 45001500 arasındaki ortak arkadaşlar:", ortaklar)
}
