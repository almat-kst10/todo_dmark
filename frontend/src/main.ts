import {createApp} from 'vue'
import App from './App.vue'
import './style.css';
import UiButton from "./components/UiComponents/UiButton.vue"
import UiInput from "./components/UiComponents/UiInput.vue"


const app = createApp(App)
app
.component('UiButton', UiButton)
.component('UiInput', UiInput)
.mount('#app')
