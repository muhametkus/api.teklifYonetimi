# SaaS Backend API

Go ile yazılmış modern bir SaaS backend uygulaması.

## 📁 Klasör Yapısı

```
saas-backend/
│
├── cmd/
│   └── server/
│       └── main.go              # Uygulama giriş noktası
│
├── internal/
│   ├── api/
│   │   ├── handlers/            # HTTP request handler'ları
│   │   │   └── ping_handler.go
│   │   ├── routes/              # Route tanımlamaları
│   │   │   └── routes.go
│   │   └── middleware/          # Middleware'ler (şimdilik boş)
│   │
│   ├── config/                  # Konfigürasyon yönetimi
│   │   └── config.go
│   │
│   ├── database/                # Veritabanı bağlantı yönetimi
│   │   └── database.go
│   │
│   ├── models/                  # Veri modelleri
│   │   └── dummy.go
│   │
│   ├── services/                # İş mantığı servisleri
│   │   └── dummy.go
│   │
│   └── utils/                   # Yardımcı fonksiyonlar
│       └── response.go          # Standart JSON response'lar
│
└── go.mod                       # Go modül tanımlamaları
```

## 🚀 Kurulum

1. Projeyi klonlayın:
```bash
cd /Users/muhammetkus/Desktop/api.teklifYonetimi
```

2. Bağımlılıkları yükleyin:
```bash
go mod download
```

3. Uygulamayı çalıştırın:
```bash
go run cmd/server/main.go
```

## 📡 API Endpoints

### Health Check
- `GET /api/v1/ping` - Server sağlık kontrolü
- `GET /api/v1/hello` - Basit merhaba endpoint'i
- `GET /hello` - Root level hello endpoint'i

### Örnek Kullanım

```bash
# Ping endpoint
curl http://localhost:8081/api/v1/ping

# Response:
# {"success":true,"message":"pong","data":{"message":"Server is running","status":"healthy"}}

# Hello endpoint
curl http://localhost:8081/hello

# Response:
# {"success":true,"message":"Hello endpoint","data":{"message":"Merhaba Go!"}}
```

## 🔧 Konfigürasyon

Uygulama aşağıdaki environment variable'ları destekler:

- `SERVER_PORT` - Server port'u (varsayılan: 8081)
- `DB_HOST` - Veritabanı host'u (varsayılan: localhost)
- `DB_PORT` - Veritabanı port'u (varsayılan: 5432)
- `DB_USER` - Veritabanı kullanıcı adı (varsayılan: postgres)
- `DB_PASSWORD` - Veritabanı şifresi
- `DB_NAME` - Veritabanı adı (varsayılan: saas_db)

## 🛠️ Teknolojiler

- **Go 1.25.5** - Programlama dili
- **Gin** - Web framework
- **PostgreSQL** - Veritabanı (opsiyonel, şimdilik devre dışı)

## 📝 Geliştirme Notları

- Database bağlantısı şimdilik devre dışı bırakılmıştır
- Models ve Services klasörleri ileride kullanılmak üzere hazırlanmıştır
- Middleware klasörü boş bırakılmıştır, gerektiğinde middleware'ler eklenebilir

## 🔜 Gelecek Özellikler

- [ ] PostgreSQL entegrasyonu
- [ ] Authentication & Authorization
- [ ] CRUD operasyonları
- [ ] Middleware'ler (CORS, Logger, vb.)
- [ ] Unit testler
- [ ] Docker desteği
