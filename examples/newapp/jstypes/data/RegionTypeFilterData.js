
	import {RegionTypeFilter, RegionTypeFilterFilter} from '../apiModel';

let regionTypeFilterData = function () {
    return {
        regionTypeFilterFilter: new RegionTypeFilterFilter(),
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
        currentRegionTypeFilterItem: {
            item: new RegionTypeFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default regionTypeFilterData;

