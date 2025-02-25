function alternateCSS () {
const styleSheetHref = document.getElementById('style_sheet').getAttribute('href');

if (styleSheetHref === '../Static/homePage.css') {
    document.getElementById('style_sheet').setAttribute("href",'../Static/alternative.css')
}
else {
    document.getElementById('style_sheet').setAttribute("href", '../Static/homePage.css')
}

}

function showFilters() {
    const filtersContainerDisplay = document.querySelector('.filters-container').getAttribute('style');

    if (filtersContainerDisplay === 'display: none;') {
        document.querySelector('.filters-container').setAttribute("style",'display: block;')
    }
    else {
        document.querySelector('.filters-container').setAttribute("style",'display: none;')
    }
}
