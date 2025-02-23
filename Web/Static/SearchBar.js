window.addEventListener('load', () => {
    // Get the JSON data
    const jsonData = document.getElementById("json-data");
    // Parse it
    const data = JSON.parse(jsonData.textContent);

    // Get the search input
    document.getElementById("searchBar").addEventListener('input', e => {
        // Get the user input
        const input = e.target.value.toLowerCase();
        console.log(`User input: ${input}`);

        // Search it in the groups names
        data.groups.forEach( (group)  => {
            if (group.name.toLowerCase().includes(input)) {
                console.log(`Group found: ${group.name}`);
            }
        });

    });
});