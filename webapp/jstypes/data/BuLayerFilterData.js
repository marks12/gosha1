
	import {BuLayerFilter, BuLayerFilterFilter} from '../apiModel';

let buLayerFilterData = function () {
    return {
        buLayerFilterFilter: new BuLayerFilterFilter(),
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
        currentBuLayerFilterItem: {
            item: new BuLayerFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default buLayerFilterData;

