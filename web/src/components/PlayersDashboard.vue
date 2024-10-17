<script setup lang="ts">

import 'datatables.net-bs5';
import DataTablesCore from 'datatables.net-dt';
import 'datatables.net-responsive-dt';
import DataTable from 'datatables.net-vue3';

DataTable.use(DataTablesCore);

const columns = [
    {
        label: 'Name',
        field: 'name',
    },
    {
        label: 'Age',
        field: 'age',
        type: 'number',
    },
    {
        label: 'Created On',
        field: 'createdAt',
        type: 'date',
        dateInputFormat: 'yyyy-MM-dd',
        dateOutputFormat: 'MMM do yy',
    },
    {
        label: 'Percent',
        field: 'score',
        type: 'percentage',
    },
]

const rows = [
    { id: 1, name: "John", age: 20, createdAt: '', score: 0.03343 },
    { id: 2, name: "Jane", age: 24, createdAt: '2011-10-31', score: 0.03343 },
    { id: 3, name: "Susan", age: 16, createdAt: '2011-10-30', score: 0.03343 },
    { id: 4, name: "Chris", age: 55, createdAt: '2011-10-11', score: 0.03343 },
    { id: 5, name: "Dan", age: 40, createdAt: '2011-10-21', score: 0.03343 },
    { id: 6, name: "John", age: 20, createdAt: '2011-10-31', score: 0.03343 },
]
</script>

<template>
    <div class="container">
        <h2><b>Players</b></h2>

        <nav aria-label="Breadcrumb">
            <ol class="breadcrumb">
                <RouterLink to="/dashboard">Dashboard</RouterLink> /
                <RouterLink to="/players">Players</RouterLink>
            </ol>
        </nav>

        <div class="table-wrapper">
            <vue-good-table :columns="columns" :rows="rows" />

        </div>
    </div>
</template>


<style scoped>
.container {
    display: flex;
    flex-direction: column;
}

.breadcrumb {
    margin-top: 1rem;
}
</style>
