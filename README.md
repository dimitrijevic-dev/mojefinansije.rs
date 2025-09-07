# 🎓 FON - Platforma za Finansijsku Pismenost

**Projekat za FON Hackathon 7.0 - Finalno kolo**

[![Go Version](https://img.shields.io/badge/Go-1.25-00ADD8?style=flat-square&logo=go)](https://golang.org/)
[![Supabase](https://img.shields.io/badge/Supabase-Database-3ECF8E?style=flat-square&logo=supabase)](https://supabase.com/)
[![OpenAI](https://img.shields.io/badge/OpenAI-GPT-412991?style=flat-square&logo=openai)](https://openai.com/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square)](https://opensource.org/licenses/MIT)

**Interaktivna platforma za učenje finansijske pismenosti sa AI asistentom, video lekcijama i personalizovanim praćenjem napretka.**

## ✨ Ključne Funkcionalnosti

- **📚 Interaktivne Lekcije** - Video sadržaj za finansijsko obrazovanje
- **🤖 AI Chat Asistent** - GPT asistent za personalizovano učenje
- **🎯 Praćenje Napretka** - Monitoring završenih lekcija i učenja
- **🔒 Bezbedna Autentifikacija** - SHA256 enkripcija lozinki
- **🔍 Pametna Pretraga** - Pronalaženje lekcija po naslovu
- **📊 Analitika** - Praćenje broja lekcija i napretka korisnika
- **🃏 Kartice za Ponavljanje** - AI-generisane kartice za ključne koncepte

## 🛠️ Tehnološki Stek

- **Backend:** Go 1.25 + Gin Framework
- **Baza podataka:** Supabase (PostgreSQL)
- **AI:** OpenAI GPT API
- **Biblioteke:** pgx/v5, scany/v2, gin-contrib/cors, openai-go

## 🏗️ Arhitektura

```
fon/
├── main.go              # Glavna tačka ulaska
├── config/              
│   ├── config.go        # Konfiguracija okruženja
│   └── encryption.go    # SHA256 enkripcija lozinki
├── router/server.go     # HTTP rute i hendleri
├── persistence/         # Operacije sa bazom podataka
│   └── database.go      # CRUD operacije za korisnike/lekcije
└── genai/              # AI integracija
    └── genai.go        # OpenAI GPT integracija
```

## 📊 Šema Baze Podataka

**Tabela Korisnici:**
- `email` (string) - Email adresa korisnika
- `displayname` (string) - Ime za prikaz
- `password` (string) - SHA256 enkriptovana lozinka

**Tabela Lekcije:**
- `id` (int8) - Identifikator lekcije
- `title` (string) - Naslov lekcije
- `body` (string) - Sadržaj lekcije
- `video_link` (string) - URL videa

**Tabela Korisnik_Lekcije:**
- `email` (string) - Email korisnika
- `lessonid` (int) - ID završene lekcije

## ⚡ Performance Karakteristike

- **Connection Pooling** - Optimizovane Supabase konekcije
- **CORS Podrška** - Cross-origin zahtevi omogućeni
- **Error Handling** - Sveobuhvatno logovanje i upravljanje greškama
- **Optimizacija Pretrage** - LIKE upiti sa pattern matchingom
- **Bezbednost Lozinki** - SHA256 enkripcija korisničkih podataka

## 📄 Licenca

MIT License - pogledajte [LICENSE](LICENSE) za detalje.

---

Built with ❤️ in Go | Powered by Supabase & OpenAI