<template>
    <a-result status="404" title="404" sub-title="Sorry, the page you visited does not exist.">
        <template #extra>
            <a-button type="primary" @click.stop="onBack">Back Home</a-button>
        </template>
    </a-result>
</template>

<script setup>
import { router } from '../router';
const onBack = () => {
    router.push("/")
    setTimeout(() => {
        router.push("/")
    }, 500);
}
</script>

<style>
</style>