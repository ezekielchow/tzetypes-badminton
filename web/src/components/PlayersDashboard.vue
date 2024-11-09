<script setup lang="ts">

import type { Player } from "@/repositories/clients/private";
import { usePlayerStore } from "@/stores/player-store";
import { useUserStore } from "@/stores/user-store";
import { reactive, ref } from "vue";
import { useRouter } from "vue-router";
import VueTableLite from "vue3-table-lite/ts";

const router = useRouter(); // Access router object

const errorMessage = ref('')
const nameSearch = ref('')

const playerStore = usePlayerStore()
playerStore.setBackendUrl(import.meta.env.VITE_PROXY_URL)

const userStore = useUserStore()
userStore.setBackendUrl(import.meta.env.VITE_PROXY_URL)

const table = reactive({
    isLoading: false,
    columns: [
        {
            label: "Name",
            field: "name",
            width: "150%",
            sortable: true,
        },
        {
            label: "Actions",
            field: "quick",
            width: "10%",
            display: function (row: Player) {
                return (
                    '<button type="button" data-id="' +
                    row.id +
                    '" class="is-rows-el edit-btn primary-button">Edit</button>'
                );
            },
        }
    ],
    rows: [] as Player[],
    totalRecordCount: 0,
    sortable: {
        order: "name",
        sort: "asc",
    },
    pageSize: 10,
    page: 1,
    offset: 0,
    order: "",
    sort: ""
});

const doSearch = async (offset: number, limit: number, order: string, sort: string) => {
    table.isLoading = true
    table.pageSize = limit
    table.offset = offset
    table.order = order
    table.sort = sort

    let page = 1
    if (offset > 0) {
        page = (offset / limit) + 1
    }

    const res = await userStore.getCurrentUser()

    if (res instanceof Error) {
        errorMessage.value = res.message
        return
    }

    errorMessage.value = ""
    table.page = page

    const listPlayersRes = await playerStore.listPlayers({
        page: table.page,
        pageSize: table.pageSize,
        ownerId: res ? res.id : undefined,
        sortArrangement: `${order}_${sort}`
    })

    if (listPlayersRes instanceof Error) {
        errorMessage.value = listPlayersRes.message
        return
    }

    errorMessage.value = ""
    table.rows = (!!listPlayersRes.players && listPlayersRes.players?.length > 0) ? listPlayersRes.players : []
    table.totalRecordCount = listPlayersRes.pagination?.totalItems ?? 0
};

const tableLoadingFinish = (elements: HTMLCollection[]) => {
    table.isLoading = false;

    router.push({
        name: 'players',
        query: { offset: table.offset, limit: table.pageSize, order: table.order, sort: table.sort },
    });

    Array.prototype.forEach.call(elements, function (element: HTMLButtonElement) {
        if (element.classList.contains("edit-btn")) {
            element.addEventListener("click", (event: MouseEvent) => {
                if (event.target) {
                    router.push({
                        name: 'players/edit',
                        params: { id: element.dataset.id }
                    })
                }
            })
        }
    });
};

const handlePageUpdate = (pageNo: Number) => {
    table.page = pageNo.valueOf()
}

doSearch(0, 10, "name", "asc")

</script>

<template>
    <div class="container">
        <h2><b>Players</b></h2>

        <div class="navigation-wrapper">
            <nav aria-label="Breadcrumb">
                <ol class="breadcrumb">
                    <RouterLink to="/dashboard">Dashboard</RouterLink> /
                    <RouterLink to="/players">Players</RouterLink>
                </ol>
            </nav>
            <RouterLink to="/players/add">
                <button class="primary-button">
                    Add Player
                </button>
            </RouterLink>
        </div>

        <div class="table-wrapper">
            <div v-if="errorMessage" class="error">
                {{ errorMessage }}
            </div>

            <VueTableLite :is-loading="table.isLoading" :columns="table.columns" :rows="table.rows"
                :total="table.totalRecordCount" :sortable="table.sortable" @do-search="doSearch"
                @is-finished="tableLoadingFinish" @get-now-page="handlePageUpdate" />
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
</style>
