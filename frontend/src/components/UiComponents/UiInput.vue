<template>
    <input 
        v-model="model" 
        :type="type" 
        :placeholder="placeholder" 
        class="input" 
        v-bind="$attrs"
    />
</template>

<script lang="ts" setup>
interface IProps {
    type: string,
    placeholder: string,
    label: string
}

const props = withDefaults(defineProps<IProps>(), {type: 'text'})

const model = defineModel();
</script>

<style scoped>

</style>