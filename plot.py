import os
import pandas as pd
import matplotlib.pyplot as plt

results_dir = 'results'
file_paths = [f for f in os.listdir(results_dir) if f.endswith('.csv')]

data_list = {os.path.splitext(file)[0]: pd.read_csv(os.path.join(results_dir, file)) for file in file_paths}
all_epochs = sorted(set.union(*[set(data['Epoch']) for data in data_list.values()]))
standard_cumulative_incentive = [305451 * (epoch + 1) for epoch in range(len(all_epochs))]
loss_rates = {file_name: {validator_id: [] for validator_id in range(666)} for file_name in data_list}

for file_name, data in data_list.items():
    data['Incentive'] = data['Head'] + data['Target'] + data['Source']
    validator_ids = data['Validator Index'].unique()[:666]
    
    cumulative_incentive = {validator_id: 0 for validator_id in validator_ids}
    
    for epoch in all_epochs:
        epoch_data = data[data['Epoch'] == epoch]
        for validator_id in validator_ids:
            incentive = epoch_data[epoch_data['Validator Index'] == validator_id]['Incentive'].values
            if len(incentive) > 0:
                cumulative_incentive[validator_id] += incentive[0]
                loss = standard_cumulative_incentive[epoch] - cumulative_incentive[validator_id]
                loss_rate = loss / standard_cumulative_incentive[epoch] if standard_cumulative_incentive[epoch] != 0 else 0
                loss_rates[file_name][validator_id].append(loss_rate)
            else:
                loss_rate = 0
                loss_rates[file_name][validator_id].append(loss_rate)
output_dir = os.path.join('results', 'figures')
os.makedirs(output_dir, exist_ok=True)

for validator_id in range(666):
    plt.figure(figsize=(10, 6))
    for file_name in data_list:
        plt.plot(all_epochs, loss_rates[file_name][validator_id], label=f'f={file_name}')
    
    plt.xlabel('Epoch')
    plt.ylabel('Incentive Loss Rate (%)')
    plt.title(f'Incentive Loss Rate for Validator {validator_id} Over Epochs')
    plt.legend(loc='lower right')
    plt.tight_layout()
    plt.savefig(os.path.join(output_dir, f'validator_{validator_id}.png'))
    plt.close()
