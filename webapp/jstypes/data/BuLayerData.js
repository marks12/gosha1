
	import {BuLayer, BuLayerFilter} from '../apiModel';

let buLayerData = function () {
    return {
        buLayerFilter: new BuLayerFilter(),
        panel: {
            type: '',
            show: false,
            create: 'create',
            edit: 'edit',
            request: 'request'
        },
        panelHeaderCreate: 'Создать',
        panelHeaderEdit: 'Изменить',
        panelSubmitButtonTextCreate: 'Создать',
        panelSubmitButtonTextEdit: 'Изменить',
        currentBuLayerItem: {
            item: new BuLayer(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default buLayerData;

