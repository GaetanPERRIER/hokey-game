<script setup>

import axios from "axios";
import {ref} from "vue";

// Récupération des éléments du formulaire
const username = ref("")
const email = ref("")
const dateOfBirth = ref("")
const password = ref("")
const confirmPassword = ref("");
const error = ref("");

async function Register(e) {
    e.preventDefault();
    error.value = "";

    if (password.value !== confirmPassword.value) {
        error.value = "Les mots de passe ne correspondent pas.";
        return;
    }

    try {
        const res = await axios.post("http://localhost:3000/api/auth/register", {
            username: username.value,
            email: email.value,
            dateOfBirth: dateOfBirth.value,
            password: password.value
        });

        // Redirection vers la page de connexion après une inscription réussie
        window.location.href = "/auth/login";
    } catch (err) {
        error.value = err.response?.data?.message || "Erreur lors de l'inscription.";
    }
}

</script>

<template>
    <form @submit="Register">
        <label for="username">Username</label>
        <input v-model="username" type="text" id="username" name="username" required />

        <label for="mail">Adresse mail</label>
        <input v-model="email" type="email" id="mail" name="mail" required />

        <label for="dateOfBirth">Date de naissance</label>
        <input v-model="dateOfBirth" type="date" id="dateOfBirth " name="dateOfBirth" required />

        <label for="password">Mot de passe</label>
        <input v-model="password" type="password" id="password" name="password" required />

        <label for="confirmPassword">Confirmer le mot de passe</label>
        <input v-model="confirmPassword" type="password" id="confirmPassword" name="confirmPassword" required />

        <p v-if="error" style="color: red;">{{ error }}</p>

        <button type="submit" :disabled="!password || !confirmPassword || password !== confirmPassword">S'inscrire</button>
    </form>
</template>

<style scoped lang="scss">

form {
    display: flex;
    flex-direction: column;
    width: 300px;
    gap: 10px;

    input {
        padding: 10px;
        font-size: 16px;
        border-radius: 5px;
        border: 1px solid black;
        width: 100%;
    }
}

</style>