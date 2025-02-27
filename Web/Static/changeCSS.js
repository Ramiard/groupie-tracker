function alternateCSS() {
    const styleSheetHref = document.getElementById('style_sheet').getAttribute('href');
    const contrastButton = document.getElementById('contrast_button');

    if (window.location.pathname === '/') {
        if (styleSheetHref === '../Static/homePage.css') {
            document.getElementById('style_sheet').setAttribute("href", '../Static/alternativeHomePage.css')
            contrastButton.innerText = 'dark_mode';
        } else {
            document.getElementById('style_sheet').setAttribute("href", '../Static/homePage.css')
            contrastButton.innerText = 'light_mode';
        }

    } else if (window.location.pathname === '/group') {
        if (styleSheetHref === '../Static/groupPage.css') {
            document.getElementById('style_sheet').setAttribute("href", '../Static/alternativeGroupPage.css')
            contrastButton.innerText = 'dark_mode';
        } else {
            document.getElementById('style_sheet').setAttribute("href", '../Static/groupPage.css')
            contrastButton.innerText = 'light_mode';
        }
    } else if (window.location.pathname === '/search') {
        if (styleSheetHref === '../Static/homePage.css') {
            document.getElementById('style_sheet').setAttribute("href", '../Static/alternativeHomePage.css')
            contrastButton.innerText = 'dark_mode';
        } else {
            document.getElementById('style_sheet').setAttribute("href", '../Static/homePage.css')
            contrastButton.innerText = 'light_mode';
        }
    }
}

function showFilters() {
    const filtersContainerDisplay = document.querySelector('.filters-container').getAttribute('style');

    if (filtersContainerDisplay === 'display: none;') {
        document.querySelector('.filters-container').setAttribute("style", 'display: block;')
    } else {
        document.querySelector('.filters-container').setAttribute("style", 'display: none;')
    }
}

function showMembers() {
    const membersListContainer = window.document.getElementById('members_list').getAttribute('style');

    if (membersListContainer === 'display: none;') {
        window.document.getElementById('members_list').setAttribute("style", 'display: block;')
        window.document.querySelector('.group-members span').innerText = '⬇️';
    } else {
        window.document.getElementById('members_list').setAttribute("style", 'display: none;')
        window.document.querySelector('.group-members span').innerText = '➡️';
    }
}

function showConcerts() {
    const concertsListContainer = window.document.getElementById('concerts_list').getAttribute('style');

    if (concertsListContainer === 'display: none;') {
        window.document.getElementById('concerts_list').setAttribute("style", 'display: block;')
        window.document.querySelector('.group-concerts span').innerText = '⬇️';
    } else {
        window.document.getElementById('concerts_list').setAttribute("style", 'display: none;')
        window.document.querySelector('.group-concerts span').innerText = '⬇️';
    }
}