window.addEventListener('load', () => {
    // Get the JSON data
    const jsonData = document.getElementById("json-data");
    // Parse it
    const data = JSON.parse(jsonData.textContent);


    // data.groups[0].members.forEach((member) => {
    //     console.log(`${member}`);
    // })

    // Get the search input
    let searchInput = document.getElementById("searchBar");
    document.getElementById("searchBar").addEventListener('input', e => {

        // Get the user input
        const input = e.target.value.toLowerCase();
        const suggestions = document.getElementById("suggestionList");
        // Clear the suggestions
        suggestions.innerHTML = "";
        console.log(`User input: ${input}`);


        if (input.length >= 2) {
            // Search it in the groups names
            data.groups.forEach((group) => {
                // Check if the group 'name' contains the input
                if (group.name.toLowerCase().includes(input)) {
                    makeSuggestion(`${group.name}`,'Group name');
                }
                // Check if the group 'members' contains the input
                group.members.forEach((member) => {
                    if (member.toLowerCase().includes(input)) {
                        makeSuggestion(`${member}`,`Member of ${group.name}`);
                    }
                });
                // Check if the group 'number of members' contains the input
                // if (group.qtyOfMembers.toString().includes(input)) {
                //     makeSuggestion(`Group found by 'QTY of members': ${group.name} | Number of members: ${group.qtyOfMembers}`);
                // }

                // Check if the group 'creation date' contains the input
                // If the input is to small we don't suggest the creation date (to avoid to flood the suggestion)
                if (input.length > 3) {
                    if (group.creationDate.toString().includes(input)) {
                        makeSuggestion(`${group.creationDate}`,`${group.name}'s creation date`);
                    }
                }
                // Check if the group 'first album' date contains the input
                // If the input is to small we don't suggest the first album date (to avoid to flood the suggestion)
                if (input.length > 3) {
                    if (group.firstAlbum.includes(input)) {
                        makeSuggestion(`${group.firstAlbum}`,`${group.name}'s first album date`);
                    }
                }
                // Check if the group 'locations' and 'dates' contains the input
                for (const [key, value] of Object.entries(group.allRelations.datesLocations)) {
                    if (key.toLowerCase().includes(input)) {
                        makeSuggestion(`${key}`,`${group.name}'s concert location`);
                    }
                    if (input.length > 3) {
                        value.forEach((date) => {
                            if (date.includes(input)) {
                                makeSuggestion(`${date}`,`${group.name}'s,  ${key} concert date`);
                            }
                        });
                    }
                }

            });
        }


    });
    function makeSuggestion(value,text) {
        const suggestion = document.createElement("option");
        suggestion.setAttribute("value", value);
        suggestion.setAttribute("class", "suggestion")
        suggestion.textContent = text;
        document.getElementById("suggestionList").appendChild(suggestion);
    }
});