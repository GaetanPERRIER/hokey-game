const Role = require("../models/Role")

const initRoles = async () => {
    const count = await Role.countDocuments()

    if (count === 0) {
        await Role.create([
            { name: "admin" },
            { name: "player" }
        ])
        console.log("Rôles créés")
    }
}

module.exports = initRoles
