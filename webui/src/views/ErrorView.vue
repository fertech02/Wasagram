<template>
    <div class="error-container">
        <h1>Error</h1>
        <p>Problem solving your request</p>
    </div>
</template>

<script>
    export default {
        name: 'ErrorView'
    }
</script>

<style scoped>
    .error-container {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100%;
    }
</style>