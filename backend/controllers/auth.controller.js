// controllers/auth.controller.js
const bcrypt = require("bcrypt")
const jwt = require("jsonwebtoken")
const User = require("../models/User")
const Role = require("../models/Role")

// Fonction pour enregistrer un nouvel utilisateur
exports.register = async (req, res) => {
    try {
        const { email, password, username, dateOfBirth } = req.body

        const existingUser = await User.findOne({ email })
        if (existingUser)
            return res.status(400).json({ message: "Email déjà utilisé" })

        const role = await Role.findOne({ name: "player" })
        const hashedPassword = await bcrypt.hash(password, 10)

        await User.create({
            role: role._id,
            username: username,
            email : email,
            dateOfBirth : dateOfBirth,
            password: hashedPassword
        })

        res.status(201).json({ message: "Utilisateur créé" })
    } catch (error) {
        res.status(500).json({ message: "Erreur serveur" })
    }
}

// Fonction pour connecter un utilisateur existant
exports.login = async (req, res) => {
    try {
        const { email, password } = req.body
        const user = await User.findOne({ email })
        if (!user)
            return res.status(400).json({ message: "Utilisateur introuvable" })

        const validPassword = await bcrypt.compare(password, user.password)
        if (!validPassword)
            return res.status(400).json({ message: "Mot de passe incorrect" })

        const token = jwt.sign(
            { userId: user._id },
            process.env.JWT_SECRET,
            { expiresIn: "1h" }
        )

        res.json({ token })
    } catch (error) {
        res.status(500).json({ message: "Erreur serveur" })
    }
}