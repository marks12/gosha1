
	import {RegionFilter, RegionFilterFilter} from '../apiModel';

let regionFilterData = function () {
    return {
        regionFilterFilter: new RegionFilterFilter(),
        panel: {
            type: '',
            show: false,
            create: 'create',
            edit: 'edit'
        },
        panelHeaderCreate: 'Создать',
        panelHeaderEdit: 'Изменить',
        panelSubmitButtonTextCreate: 'Создать',
        panelSubmitButtonTextEdit: 'Изменить',
        currentRegionFilterItem: {
            item: new RegionFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default regionFilterData;

