# 🏒 Hockey Game – MVP Roadmap

Objectif : obtenir un **jeu multijoueur jouable** (MVP), simple mais solide, avec :
- Backend **Go authoritative**
- Frontend **Vue.js**
- Communication **WebSocket**

---

## 🎯 Vision du MVP

- 2 joueurs dans un match
- Vue top-down 2D
- Déplacements basiques
- Un puck qui bouge
- Score simple
- Le serveur décide de tout

---

## 🧱 Étape 0 — Setup & fondations

### Backend
- [ ] Installer Go
- [ ] `go mod init`
- [ ] Structure de projet propre (`cmd/`, `internal/`)
- [ ] Serveur HTTP Go fonctionnel
- [ ] Endpoint `/health`

### Frontend
- [ ] Projet Vue 3 initialisé
- [ ] App qui démarre sans erreur

🎯 Résultat : le projet démarre côté front et back

---

## 🔌 Étape 1 — WebSocket fonctionnel

### Backend
- [ ] Ajouter Gorilla WebSocket
- [ ] Endpoint `/ws`
- [ ] Connexion / déconnexion joueur
- [ ] Logs clairs côté serveur

### Frontend
- [ ] Connexion WebSocket
- [ ] Réception de messages serveur
- [ ] Envoi de messages simples

🎯 Résultat : le client parle au serveur en temps réel

---

## 👥 Étape 2 — Match / Room simple

### Backend
- [ ] Structure `Match`
- [ ] Max 2 joueurs par match
- [ ] Join / Leave
- [ ] État du match (waiting / playing)

🎯 Résultat : 2 joueurs peuvent rejoindre le même match

---

## ⏱️ Étape 3 — Game Loop serveur (cœur du jeu)

### Backend
- [ ] Tick serveur (30 ou 60 Hz)
- [ ] Boucle indépendante des clients
- [ ] Broadcast régulier de l’état

🎯 Résultat : le serveur vit tout seul

---

## 🧍 Étape 4 — Joueurs & inputs

### Backend
- [ ] Structure `Player`
- [ ] Inputs = intentions (haut, bas, gauche, droite)
- [ ] Stockage des inputs
- [ ] Application dans la game loop

### Frontend
- [ ] Capturer clavier
- [ ] Envoyer inputs au serveur

🎯 Résultat : les joueurs peuvent bouger

---

## 🏒 Étape 5 — Puck & physique minimale

### Backend
- [ ] Structure `Puck`
- [ ] Mouvement simple
- [ ] Rebonds sur les murs
- [ ] Terrain avec limites

🎯 Résultat : ça commence à ressembler à du hockey

---

## 📡 Étape 6 — Synchronisation état du jeu

### Backend
- [ ] État global du match
- [ ] Envoi snapshots réguliers

### Frontend
- [ ] Réception état
- [ ] Affichage via canvas
- [ ] Interpolation simple

🎯 Résultat : le jeu est visible et fluide

---

## 🥅 Étape 7 — Règles de base

### Backend
- [ ] Détection de but
- [ ] Score
- [ ] Reset puck
- [ ] Fin de match simple

🎯 Résultat : un match complet jouable

---

## 🚀 Étape 8 — Polish MVP (optionnel)

- [ ] Affichage score
- [ ] Indication joueur local
- [ ] Reconnexion simple
- [ ] Logs propres
- [ ] README clair

---

## ❌ Hors scope MVP (à ne PAS faire maintenant)

- Auth avancée
- Skins / animations
- Chat
- DB complexe
- Matchmaking automatique
- Anti-cheat avancé

---

## ✅ Définition de "MVP terminé"

- [ ] 2 joueurs peuvent jouer un match
- [ ] Le serveur est authoritative
- [ ] Le jeu est stable 5–10 minutes
- [ ] Pas de crash serveur
- [ ] Code lisible et structuré

---

## 🧠 Next steps (post-MVP)

- Spectateurs
- Classement
- Matchmaking
- Replays
- Scaling (Redis, multi serveurs)

---

🎉 **Si tu arrives ici : tu as un vrai jeu multijoueur.**
