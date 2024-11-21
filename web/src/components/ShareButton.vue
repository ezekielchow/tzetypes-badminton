<script setup lang="ts">
import shareImg from '@/assets/images/share.png';

const props = defineProps({
    title: String,
    text: String,
    url: String
})

const shareContent = async () => {
    if (navigator.share) {
        try {
            await navigator.share({
                title: props.title,
                text: props.text,
                url: props.url,
            });
            console.log("Content shared successfully!");
        } catch (err) {
            console.error("Error sharing:", err);
        }
    } else {
        // Fallback: Open WhatsApp with pre-filled message
        const message = `${props.text} \nCheck it out here: ${props.url}`;
        window.open(`https://wa.me/?text=${encodeURIComponent(message)}`, "_blank");
    }
};
</script>

<template>
    <img @click="shareContent" :src="shareImg" width="25px" height="25px">
</template>

<style scoped></style>