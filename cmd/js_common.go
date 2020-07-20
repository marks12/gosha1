package cmd

const commonContent = `

export function findItemIndex(list, fn) {
    let index;
    
    if (!list || !fn) return;
    
    for (let i = 0; i < list.length; i++) {
        
        if (fn(list[i])) {
            index = i;
            break;
        }
        
    }
    return index;
}

export function transformType(data, entity) {

    if (Array.isArray && Array.isArray(data)) {
        data.forEach(item => {
            for (let key in item) {
                if (entity.hasOwnProperty(key) && Number.isInteger(entity[key])) {
                    item[key] = Number(item[key])
                }
            }
        });
        return data;
    }

    for (let key in data) {
        if (entity.hasOwnProperty(key) && Number.isInteger(entity[key])) {
            data[key] = Number(data[key])
        }
    }
    return data;
}




`