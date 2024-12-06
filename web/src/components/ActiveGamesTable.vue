<template>
    <div class="table-container">
        <div v-if="errorMessage" class="error">
            {{ errorMessage }}
        </div>
        <table class="responsive-table">
            <thead>
                <tr>
                    <th>Left Player(s)</th>
                    <th>Right Player(s)</th>
                    <th>Started At</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="game in games" :key="game.id" @click="goToGame(game.id)" class="table-row">
                    <td>{{ game.leftPlayers }}</td>
                    <td>{{ game.rightPlayers }}</td>
                    <td>{{ game.startedAt }}</td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script lang="ts">
import { GameTypes } from "@/enums/game";
import { useGameStore } from "@/stores/game-store";
import { DateTime } from 'luxon';
import { defineComponent, ref } from "vue";
import { useRouter } from "vue-router";

export default defineComponent({
    name: "GameTable",
    async setup() {
        const errorMessage = ref("")
        const router = useRouter();
        const gameStore = useGameStore()
        gameStore.setBackendUrl(import.meta.env.VITE_PROXY_URL)

        const listActiveGames = async () => {
            const games = ref([] as any);

            const activeGames = await gameStore.listActiveGames()

            if (activeGames instanceof Error) {
                errorMessage.value = activeGames.message
                return
            }

            if (!!activeGames.games && activeGames.games?.length > 0) {
                for (let i = 0; i < activeGames.games.length; i++) {
                    const game = activeGames.games[i];

                    const startedAtDate = DateTime.fromFormat(
                        game.createdAt.split(".")[0],
                        'yyyy-MM-dd HH:mm:ss',
                        { zone: 'utc' } // Assume it's in UTC
                    );
                    const startedAt = startedAtDate.toLocaleString(DateTime.TIME_SIMPLE);

                    games.value.push({
                        id: game.id,
                        leftPlayers: game.gameType == GameTypes.GAME_TYPE_SINGLES ? game.leftEvenPlayerName : `${game.leftOddPlayerName}, \n${game.leftEvenPlayerName}`,
                        rightPlayers: game.gameType == GameTypes.GAME_TYPE_SINGLES ? game.rightEvenPlayerName : `${game.rightEvenPlayerName}, \n${game.rightOddPlayerName}`,
                        startedAt: startedAt
                    })
                }
            }

            return games
        }

        const games = await listActiveGames()

        // Redirect to ongoing game page
        const goToGame = (id: number) => {
            router.push({
                name: "game/playing",
                params: {
                    id,
                }
            });
        };

        return {
            games,
            goToGame,
            errorMessage
        };
    },
});
</script>

<style scoped>
.table-container {
    overflow-x: auto;
}

.responsive-table {
    width: 100%;
    border-collapse: collapse;
    font-size: 14px;
    text-align: left;
}

thead {
    background-color: #343a40;
    color: #ffffff;
}

thead th {
    padding: 10px;
}

tbody tr {
    cursor: pointer;
    /* Changes cursor to pointer on hover */
    transition: background-color 0.3s ease, box-shadow 0.3s ease;
    /* Smooth transition */
}

tbody tr:hover {
    background-color: #e9ecef;
    /* Light gray for hover */
    box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
    /* Subtle shadow for a clickable effect */
}

tbody tr:active {
    background-color: #d6d8db;
    /* Slightly darker gray on click/tap */
}

tbody td {
    padding: 12px;
    /* Increased padding for easier tap on mobile */
    border-bottom: 1px solid #ddd;
}

/* Mobile styles */
@media (max-width: 768px) {
    .responsive-table {
        font-size: 12px;
    }

    thead th,
    tbody td {
        padding: 14px;
        /* Further increased padding for better touch targets */
    }
}
</style>
