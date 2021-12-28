function getBanksFromServer() {
    console.log(">> getBanksFromServer");
    fetch('//localhost:9090/banks', {
        method: 'GET'
    }).then(respnose => respnose.json())
        .then(respnose => {
            console.log('Response', respnose)
            showListOfBanks(respnose);
        }).catch(error => console.error('Error', error))
}

let bankCodeSelector = document.getElementById("selector4banks");
bankCodeSelector.addEventListener("change", () => {
    console.log("Bank was changed to", bankCodeSelector.value);
    queryBranchesByConditions();
})

let provinceSelector = document.getElementById("selector4provinces");
provinceSelector.addEventListener("change", () => {
    console.log("Province was changed to", provinceSelector.value);
    queryBranchesByConditions();
})

let citySelector = document.getElementById("selector4cities");
citySelector.addEventListener("change", () => {
    console.log("City was changed to", citySelector.value);
    queryBranchesByConditions();
})



function showListOfBranches(branchArray) {
    console.log(">> showListOfBranches", branchArray);

    let result = "";
    for (let i = 0; i < branchArray.length; i++) {
        let branchObj = branchArray[i];
        result += `${branchObj.branchCode}:${branchObj.name}<br/>`;
    }

    document.getElementById("result").innerHTML = result;
}

function cleanResult() {
    console.log(">> cleanResult");
    let result = document.getElementById("result");
    result.innerText = "";
    showAlert(null);
    document.getElementById("number_of_result").innerText = "0";
}

function showAlert(errorMessage) {
    console.log(">> showAlert:", errorMessage);

    let msgbar = document.getElementById("alert");

    if (errorMessage == null) {
        msgbar.style.visibility = "hidden"
    } else {
        msgbar.innerText = errorMessage;
        msgbar.style.visibility = "visible"
    }
}

function showLoadingStatus(isLoading) {
    let lodaingBar = document.getElementById("loading_bar");
    lodaingBar.style.visibility = (isLoading) ? "visible" : "hidden";
}

function showListOfBanks(bankObjArray) {
    console.log(">> showListOfBanks", bankObjArray);

    let defaultOfListItemView = document.getElementById("bank_list_item_template");

    for (let i = 0; i < bankObjArray.length; i++) {
        let bankObj = bankObjArray[i];

        let listItem = defaultOfListItemView.cloneNode(true);
        listItem.innerText = `${bankObj.code}:${bankObj.name}`;
        listItem.removeAttribute('selected');
        listItem.value = `${bankObj.code}`;
        bankCodeSelector.appendChild(listItem);
    }
}

function queryBranchesByConditions() {

    let bankCode = bankCodeSelector.value;
    let province = provinceSelector.value;
    let city = citySelector.value;
    console.log(">> queryBranchesByConditions:", bankCode, province, city);

    let url = "//localhost:9090/branches";
    let query = "";
    if (bankCode != "---") {
        query = `${query}bankcode=${bankCode}`;
    }
    if (province != "---") {
        if (query.length > 0) {
            query += "&";
        }
        query = `${query}province=${province}`;
    }
    if (city != "---") {
        if (query.length > 0) {
            query += "&";
        }
        query = `${query}city=${city}`;
    }
    if (query.length > 0) {
        url = `${url}?${query}`;
    }

    console.log("Query...");
    console.log(url);
    let url4API = document.getElementById("target_url");
    url4API.innerText = url;
    url4API.href = url;

    cleanResult();
    showLoadingStatus(true);
    fetch(url, {
        method: 'GET'
    }).then(respnose => {
        console.log("status code:", respnose.status); // returns 200
        return respnose.json()
    }).then(response => {

        showLoadingStatus(false);

        if (response.length == null) {
            showAlert(response.message)
            return
        }
        console.log('Response', response)
        showListOfBranches(response);
        document.getElementById("number_of_result").innerText = response.length ? response.length : "0";
        showAlert(null);
    })
        .catch(function (error) {
            console.log("Error:", error);
        })
}

getBanksFromServer();