window.addEventListener('load', () => {
    // Get the JSON data
    const jsonData = document.getElementById("json_data");
    // Parse it
    const data = JSON.parse(jsonData.textContent);

    // Create the suggestions container
    const suggestions =  document.createElement("ul");
    suggestions.setAttribute("id","suggestions_container");

    // 'i' will be used to know when to put the suggestion_container on the html
    let i = 0;

    // Get the search input
    document.getElementById("search_bar").addEventListener('input', e => {
        // Get the user input
        const input = e.target.value.toLowerCase();

        i = input.length;
        // Check the 'i' value to know when to put the suggestions container on the html
        // if the user didn't write anything we don't need to show the suggestions
        if (i === 1) {
            e.target.insertAdjacentElement('afterend',suggestions);
        } else if (i === 0) {
            document.getElementById("suggestions_container").remove();
        }

        // Clear the suggestions at each input to know show the old suggestions
        suggestions.innerHTML = "";
        console.log(`User input: ${input}`);


        if (input.length >= 1) {
            // Parcours all the groups
            data.groups.forEach((group) => {
                // Check if the group 'name' contains the input
                if (group.name.toLowerCase().includes(input)) {
                    makeSuggestion(`${group.image}`,`${group.name}`,`/`);
                }

                // Check if the group 'members' contains the input
                group.members.forEach((member) => {
                    if (member.toLowerCase().includes(input)) {
                        makeSuggestion(`${group.image}`,`|${member}| Member of ${group.name}`,`/`);
                    }
                });

                // Check if the group 'creation date' contains the input
                // If the input is to small we don't suggest the creation date (to avoid to flood the suggestion)
                if (input.length > 3) {
                    if (group.creationDate.toString().includes(input)) {
                        makeSuggestion(`${group.image}`,`|${group.creationDate}| ${group.name}'s creation date`,`/`);
                    }
                }

                // Check if the group 'first album' date contains the input
                // If the input is to small we don't suggest the first album date (to avoid to flood the suggestion)
                if (input.length > 3) {
                    if (group.firstAlbum.includes(input)) {
                        makeSuggestion(`${group.image}`,`|${group.firstAlbum}| ${group.name}'s first album date`,`/`);
                    }
                }

                // Check if the group 'locations' and 'dates' contains the input
                for (const [key, value] of Object.entries(group.allRelations.datesLocations)) {
                    if (key.toLowerCase().includes(input)) {
                        makeSuggestion(`${group.image}`,`|${key}| ${group.name}'s concert location`,`/`);
                    }
                    if (input.length > 3) {
                        value.forEach((date) => {
                            if (date.includes(input)) {
                                makeSuggestion(`${group.image}`,`|${date}| ${group.name}'s  ${key} concert date`,`/`);
                            }
                        });
                    }
                }

            });
        }


    });
    // 'makeSuggestion' will create a suggestion and append it to the suggestions container
    function makeSuggestion(img,text,url) {
        // Make the suggestion base and append it to the suggestions container
        const suggestion = document.createElement("li");
        suggestion.setAttribute("class","suggestion")
        document.getElementById("suggestions_container").appendChild(suggestion);

        // Make the suggestion image and append it to the suggestion
        const suggestionImg = document.createElement("img");
        suggestionImg.setAttribute("class","suggestion-image");
        suggestionImg.setAttribute("src", img);
        suggestion.appendChild(suggestionImg);

        // Make the suggestion text and append it to the suggestion img
        const suggestionText = document.createElement("a");
        suggestionText.setAttribute("class","suggestion-text");
        suggestionText.setAttribute("href",url);
        suggestionText.textContent = text;
        suggestionImg.insertAdjacentElement('afterend',suggestionText);
    }
});