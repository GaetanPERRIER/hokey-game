<script setup>

import axios from "axios";
import {ref} from "vue";

// Récupération des éléments du formulaire
const email = ref("")
const password = ref("")

async function Login(e) {
    e.preventDefault()

    try {
        const res = await axios.post("http://localhost:3000/api/auth/login", {
            email: email.value,
            password: password.value
        })

        // Stockage du token dans le localStorage
        localStorage.setItem("token", res.data.token)

        // Redirection vers la page d'accueil après une connexion réussie
        window.location.href = "/";
    } catch (err) {
    }


}

</script>

<template>
    <form @submit="Login">
        <label for="mail">Adresse mail</label>
        <input v-model="email" type="email" id="mail" name="mail" required />

        <label for="password">Mot de passe</label>
        <input v-model="password" type="password" id="password" name="password" required />

        <button type="submit">Se connecter</button>
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