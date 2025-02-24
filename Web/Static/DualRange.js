window.addEventListener('load', () =>
    document.querySelectorAll(".dualRange-container").forEach(dual_range => {
    // Get range pickers
    let ranges = dual_range.querySelectorAll("input[type=range]");
    let dualMinValue = dual_range.querySelector(".dualRange-minValue");
    let dualMaxValue = dual_range.querySelector(".dualRange-maxValue");

    // 'min' can't be higher than 'max'
    ranges[0].addEventListener('input',()=> {
        if (Number(ranges[0].value) >= Number(ranges[1].value)) {
            ranges[0].value = Number(ranges[1].value) - 1 ;
        }
        dualMinValue.innerHTML = "From : " + ranges[0].value;
    });

    // 'max' can't be lower than 'min'
    ranges[1].addEventListener('input', e=> {
        if (Number(ranges[1].value) <= Number(ranges[0].value)) {
            ranges[1].value = Number(ranges[0].value) + 1;
        }
        dualMaxValue.innerHTML = "To : "+ ranges[1].value;
    });

}));
