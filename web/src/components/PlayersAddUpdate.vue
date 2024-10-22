<script setup lang="ts">
import router from '@/router';
import { usePlayerStore } from '@/stores/player-store';
import { ref } from 'vue';
import { useRoute } from 'vue-router';

interface Props {
    isEdit: boolean;
}

const props = defineProps<Props>();
const errorMessage = ref("")
const successMessage = ref("")
const formIsLoading = ref(false)
const playerName = ref("")
const route = useRoute()

const playerStore = usePlayerStore()
playerStore.setBackendUrl(import.meta.env.VITE_BACKEND_URL)

const addPlayer = async (e: MouseEvent) => {
    formIsLoading.value = true

    console.log('did come in ');


    const res = await playerStore.addPlayer({
        addPlayerRequestSchema: {
            name: playerName.value
        }
    })

    if (res instanceof Error) {
        formIsLoading.value = false
        errorMessage.value = res.message
        return
    }

    formIsLoading.value = false

    if (e.target) {
        const target = e.target as HTMLButtonElement

        if (target.name == "submit_and_new") {
            window.location.reload()
        } else {

            router.push({
                name: 'players/edit',
                params: { id: res.id }
            })
        }
    }
}

const updatePlayer = async () => {
    if (route.params.id && typeof route.params.id == "string") {
        formIsLoading.value = true

        const res = await playerStore.updatePlayer({
            id: route.params.id,
            updatePlayerWithIdRequest: {
                name: playerName.value
            }
        })

        if (res instanceof Error) {
            formIsLoading.value = false
            errorMessage.value = res.message
            return
        }

        formIsLoading.value = false
        successMessage.value = "Successfully updated"
    }
}

const getOld = async () => {
    if (route.params.id && typeof route.params.id == "string") {
        formIsLoading.value = true

        const res = await playerStore.getPlayer({
            id: route.params.id
        })

        if (res instanceof Error) {
            formIsLoading.value = false
            errorMessage.value = res.message
            return
        }

        formIsLoading.value = false
        playerName.value = res.name
    }
}

if (props.isEdit) {
    await getOld()
}

</script>

<template>
    <div class="container">
        <h2><b>{{ props.isEdit ? "Edit Player" : "Add Player" }}</b></h2>

        <div class="navigation-wrapper">
            <nav aria-label="Breadcrumb">
                <ol class="breadcrumb">
                    <RouterLink to="/dashboard">Dashboard</RouterLink> /
                    <RouterLink to="/players">Players</RouterLink> /
                    <RouterLink v-if="props.isEdit" :to="{ name: 'players/edit', params: { id: route.params.id } }">{{
                        route.params.id }}
                    </RouterLink>
                    <RouterLink v-else to="/players/add">Add Player</RouterLink>
                </ol>
            </nav>
        </div>

        <div class="form-container">
            <p class="error-message" id="error-message" v-if='errorMessage !== ""'>{{ errorMessage }}</p>
            <p class="success-message" id="success-message" v-if='successMessage !== ""'>{{ successMessage }}</p>
            <form @submit.prevent="(event) => { event.preventDefault() }">
                <div class="form-group">
                    <label for="name">Name</label>
                    <input type="text" id="name" name="name" placeholder="Enter player name" v-model="playerName"
                        required>
                </div>
                <button v-if="props.isEdit" type="button" class="submit-btn" :disabled="formIsLoading"
                    name="submit_normal" @click="updatePlayer">Update Player</button>
                <button v-else type="button" class="submit-btn" :disabled="formIsLoading" name="submit_normal"
                    @click="addPlayer">Add Player</button>

                <button v-if="!props.isEdit" @click="addPlayer" type="button" class="submit-btn add-and-new-button"
                    :disabled="formIsLoading" name="submit_and_new">Add and New
                    Player</button>
            </form>
        </div>
    </div>
</template>


<style scoped>
.container {
    display: flex;
    flex-direction: column;
    padding: 1rem;
}

.navigation-wrapper {
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-radius: 0.25rem;
    padding: 0.5rem;
    border: 0.5px solid silver;
}

.form-container {
    margin-top: 1rem;
    align-self: center;
}

.add-and-new-button {
    margin-top: 1.5rem;
}
</style>
