// ЖС функция для генерации bash скрипта, который:
// 1. пересоздаст папку downloads_from_wgeter
// 2. Скачает файлы по ксс селектору и атрибуту
// 3. Назовёт их нормально и сохранит расширение
// 4. Если вместо url подан base64, сформирует скрипт для перевода в картинку
// Работает в chromium и firefox, лучше chromium

function wgeter(css_selector, attr){
    let to_return = "cd ~/Downloads && rm -rf downloads_from_wgeter && mkdir downloads_from_wgeter && cd downloads_from_wgeter ";
    let urls = document.querySelectorAll(css_selector);
    for(i = 0; i < urls.length; i++){
        let file_name = urls[i][attr];

        if (file_name.slice(0, 4) === "data"){
            let extension = file_name.split('/', 2)[1].split(';', 2)[0]
            let base64_data = file_name.split(',', 2)[1]
            to_return += `; echo "${base64_data}" | base64 -d > wgeter_${i}.${extension}`
        }else{
            if (file_name !== ""){
                if (file_name[0]==="/"){
                    file_name = window.location.hostname + file_name
                }
                let splited_name = file_name.split('.');
                let extension = splited_name[splited_name.length - 1].split('?')[0]
                allowed_extension = ["jpeg", "jpg", "png", "mpeg", "wbem"]
                if (!allowed_extension.includes(extension)){
                    extension = "wbem"
                }
                to_return += `;  wget -nv "${file_name}" -O "wgeter_${i}.${extension}" `;
            }
        }
    }
    console.log(to_return);
}
// https://www.123rf.com/photo_125721283_ground-parking-cars-after-snowfall-in-winter-view-from-above-top-or-aerial-view-automobiles-covered-.html
// wgeter(".imgSrcClass", "src")
