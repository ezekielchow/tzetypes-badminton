<template>
    <button :class="['button', buttonTypeClass]" :disabled="isLoading" @click="handleClick">
        <template v-if="isLoading">
            <span class="spinner"></span>
            <span class="loading-text">Loading...</span>
        </template>
        <template v-else>
            <slot />
        </template>
    </button>
</template>

<script>
import { computed } from "vue";

export default {
    name: "ButtonComponent",
    props: {
        type: {
            type: String,
            default: "primary", // 'primary' or 'secondary'
            validator: (value) => ["primary", "secondary"].includes(value),
        },
        isLoading: {
            type: Boolean,
            default: false,
        },
    },
    setup(props, { emit }) {
        const buttonTypeClass = computed(() => `button-${props.type}`);

        const handleClick = (event) => {
            if (!props.isLoading) {
                emit("click", event);
            }
        };

        return {
            buttonTypeClass,
            handleClick,
        };
    },
};
</script>

<style scoped>
.spinner {
    display: inline-block;
    width: 12px;
    height: 12px;
    border: 2px solid rgba(255, 255, 255, 0.3);
    border-top-color: #fff;
    border-radius: 50%;
    animation: spin 1s linear infinite;
    margin-right: 5px;
    vertical-align: middle;
}

.loading-text {
    vertical-align: middle;
}

@keyframes spin {
    to {
        transform: rotate(360deg);
    }
}

button:disabled {
    cursor: not-allowed;
    opacity: 0.6;
}
</style>