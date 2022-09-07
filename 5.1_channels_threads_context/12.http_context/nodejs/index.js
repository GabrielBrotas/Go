import express from "express";

const app = express()

app.get('/', async (req, res) => {
    // if the client close the request this function will still be executed,
    // and it will use the CPU/memory resources for a query that will not be used because the client closed the connection
    async function long_execution_query_to_db(ms) {
        return new Promise(resolve => {
            setTimeout(() => {
                console.log("resolved")
                resolve()
            }, ms)
        })
    }
    console.log("processing...")
    await long_execution_query_to_db(5000)

    return res.status(200).json({success: true})
})

const server = app.listen(8080, () => {
    console.log("app running on port 8080")
})

