<script setup>

</script>

<template>
    <div class="play-form">
        <router-link to="/game">Jouer</router-link>
    </div>
</template>

<style scoped lang="scss">

.play-form {
    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);

    a {
        padding: 10px 25px;
        font-size: 16px;
        border-radius: 5px;
        border: 1px solid black;
        color: black;
        text-decoration: none;
    }
}



</style>