<template>
    <button
        class="btn" 
        :class="[`btn__${type}`]"
        v-bind="$attrs"
    >
        <slot />
    </button>
</template>

<script lang="ts" setup>
type typeButton = 'primary' | 'secondary' | 'deleted'

interface IProps {
    type: typeButton
}
const props = withDefaults(defineProps<IProps>(), {type: "primary"})
</script>

<style scoped>
.btn {
    cursor: pointer;
    color: white !important;
    border: none;
    border-radius: 16px;
    font-size: 14px;
    font-weight: 400 ;
    transition: all 0.3s ease;
    padding: 6px 10px;
    &:hover {
        transform: scale(1.01);
    }
}
.btn__primary {
    background: #007bff;
}
.btn__secondary {
    background: green;
}
.btn__deleted {
    background: red;
}
</style>